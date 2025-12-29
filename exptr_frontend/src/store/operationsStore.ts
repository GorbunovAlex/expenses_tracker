// =============================================================================
// Operations Store
// =============================================================================

import { create } from "zustand";
import type { Operation } from "@/types/api";

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

interface OperationsState {
  // State
  operations: Operation[];

  // Actions
  setOperations: (operations: Operation[]) => void;
  addOperation: (operation: Operation) => void;
  updateOperation: (id: string, updates: Partial<Operation>) => void;
  removeOperation: (id: string) => void;
  clearOperations: () => void;

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

  // Actions
  setOperations: (operations) => set({ operations }),

  addOperation: (operation) =>
    set((state) => ({
      operations: [operation, ...state.operations],
    })),

  updateOperation: (id, updates) =>
    set((state) => ({
      operations: state.operations.map((op) =>
        op.id === id ? { ...op, ...updates } : op,
      ),
    })),

  removeOperation: (id) =>
    set((state) => ({
      operations: state.operations.filter((op) => op.id !== id),
    })),

  clearOperations: () => set({ operations: [] }),

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
export const useOperationsActions = () =>
  useOperationsStore((state) => ({
    setOperations: state.setOperations,
    addOperation: state.addOperation,
    updateOperation: state.updateOperation,
    removeOperation: state.removeOperation,
    clearOperations: state.clearOperations,
  }));
