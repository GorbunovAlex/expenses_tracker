// =============================================================================
// Operations Store
// =============================================================================

import { create } from "zustand";
import { useShallow } from "zustand/react/shallow";
import { operationsApi, type OperationRequest } from "@/api/client";
import type { Operation } from "@/types/api";

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

interface OperationsState {
  // State
  operations: Operation[];
  isLoading: boolean;
  error: string | null;
  _isFetching: boolean; // Internal flag for deduplication

  // Async Actions (API)
  fetchOperations: (options?: { force?: boolean }) => Promise<void>;
  createOperation: (data: OperationRequest) => Promise<void>;
  updateOperation: (
    id: string,
    data: Partial<OperationRequest>,
  ) => Promise<void>;
  deleteOperation: (id: string) => Promise<void>;

  // Sync Actions
  setOperations: (operations: Operation[]) => void;
  addOperation: (operation: Operation) => void;
  removeOperation: (id: string) => void;
  clearOperations: () => void;
  clearError: () => void;

  // Selectors / Computed
  getOperationsByType: (type: "income" | "expense") => Operation[];
  getOperationsByCategory: (categoryId: string) => Operation[];
  getOperationById: (id: string) => Operation | undefined;
  getTotalIncome: () => number;
  getTotalExpense: () => number;
  getBalance: () => number;
}

// -----------------------------------------------------------------------------
// Store
// -----------------------------------------------------------------------------

export const useOperationsStore = create<OperationsState>()((set, get) => ({
  // Initial State
  operations: [],
  isLoading: false,
  error: null,
  _isFetching: false,

  // Async Actions (API)
  fetchOperations: async (options?: { force?: boolean }) => {
    const force = options?.force ?? false;

    // Deduplicate: if already fetching, don't start another request
    if (get()._isFetching) {
      return;
    }

    // Skip if we already have operations (unless force is true)
    if (!force && get().operations.length > 0) {
      return;
    }

    set({ isLoading: true, error: null, _isFetching: true });
    try {
      const response = await operationsApi.getAll();
      if (response.operations) {
        set({ operations: response.operations });
      }
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Failed to fetch operations";
      set({ error: message });
    } finally {
      set({ isLoading: false, _isFetching: false });
    }
  },

  createOperation: async (data: OperationRequest) => {
    set({ isLoading: true, error: null });
    try {
      const response = await operationsApi.create(data);
      if (response.status === "OK" || response.status === "success") {
        // Refetch operations to get the new one with ID
        // Reset _isFetching to allow refetch (keep existing data to avoid flicker)
        set({ _isFetching: false });
        const opsResponse = await operationsApi.getAll();
        if (opsResponse.operations) {
          set({ operations: opsResponse.operations });
        }
      }
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Failed to create operation";
      set({ error: message });
      throw err;
    } finally {
      set({ isLoading: false });
    }
  },

  updateOperation: async (id: string, data: Partial<OperationRequest>) => {
    set({ isLoading: true, error: null });
    try {
      await operationsApi.update(id, data);
      // Refetch operations to get updated data
      // Reset _isFetching to allow refetch (keep existing data to avoid flicker)
      set({ _isFetching: false });
      const opsResponse = await operationsApi.getAll();
      if (opsResponse.operations) {
        set({ operations: opsResponse.operations });
      }
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Failed to update operation";
      set({ error: message });
      throw err;
    } finally {
      set({ isLoading: false });
    }
  },

  deleteOperation: async (id: string) => {
    set({ isLoading: true, error: null });
    try {
      await operationsApi.delete(id);
      // Remove from local state immediately
      set((state) => ({
        operations: state.operations.filter((op) => op.id !== id),
      }));
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Failed to delete operation";
      set({ error: message });
      throw err;
    } finally {
      set({ isLoading: false });
    }
  },

  // Sync Actions
  setOperations: (operations) => set({ operations }),

  addOperation: (operation) =>
    set((state) => ({
      operations: [operation, ...state.operations],
    })),

  removeOperation: (id) =>
    set((state) => ({
      operations: state.operations.filter((op) => op.id !== id),
    })),

  clearOperations: () => set({ operations: [], _isFetching: false }),

  clearError: () => set({ error: null }),

  // Selectors / Computed
  getOperationsByType: (type) => {
    const { operations } = get();
    return operations.filter((op) => op.type === type);
  },

  getOperationsByCategory: (categoryId) => {
    const { operations } = get();
    return operations.filter((op) => op.category_id === categoryId);
  },

  getOperationById: (id) => {
    const { operations } = get();
    return operations.find((op) => op.id === id);
  },

  getTotalIncome: () => {
    const { operations } = get();
    return operations
      .filter((op) => op.type === "income")
      .reduce((sum, op) => sum + op.amount, 0);
  },

  getTotalExpense: () => {
    const { operations } = get();
    return operations
      .filter((op) => op.type === "expense")
      .reduce((sum, op) => sum + op.amount, 0);
  },

  getBalance: () => {
    const { getTotalIncome, getTotalExpense } = get();
    return getTotalIncome() - getTotalExpense();
  },
}));

// -----------------------------------------------------------------------------
// Selector Hooks (for optimized re-renders)
// -----------------------------------------------------------------------------

export const useOperations = () =>
  useOperationsStore((state) => state.operations);

export const useOperationsLoading = () =>
  useOperationsStore((state) => state.isLoading);

export const useOperationsError = () =>
  useOperationsStore((state) => state.error);

// Use shallow equality to prevent unnecessary re-renders when actions haven't changed
export const useOperationsActions = () =>
  useOperationsStore(
    useShallow((state) => ({
      fetchOperations: state.fetchOperations,
      createOperation: state.createOperation,
      updateOperation: state.updateOperation,
      deleteOperation: state.deleteOperation,
      setOperations: state.setOperations,
      addOperation: state.addOperation,
      removeOperation: state.removeOperation,
      clearOperations: state.clearOperations,
      clearError: state.clearError,
    })),
  );
