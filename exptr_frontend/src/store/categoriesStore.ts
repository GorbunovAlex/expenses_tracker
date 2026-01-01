import { create } from "zustand";
import { useShallow } from "zustand/react/shallow";
import { categoriesApi } from "@/api/client";
import type { Category } from "@/types/api";

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

interface CategoriesState {
  // State
  categories: Category[];
  isLoading: boolean;
  error: string | null;
  _isFetching: boolean; // Internal flag for deduplication

  // Async Actions (API)
  fetchCategories: (options?: { force?: boolean }) => Promise<void>;

  // Sync Actions
  setCategories: (categories: Category[]) => void;
  addCategory: (category: Category) => void;
  updateCategory: (id: string, updates: Partial<Category>) => void;
  removeCategory: (id: string) => void;
  clearCategories: () => void;
  clearError: () => void;

  // Selectors
  getCategoryById: (id: string) => Category | undefined;
  getCategoriesByType: (type: "income" | "expense") => Category[];
}

// -----------------------------------------------------------------------------
// Store
// -----------------------------------------------------------------------------

export const useCategoriesStore = create<CategoriesState>()((set, get) => ({
  // Initial State
  categories: [],
  isLoading: false,
  error: null,
  _isFetching: false,

  // Async Actions (API)
  fetchCategories: async (options?: { force?: boolean }) => {
    const force = options?.force ?? false;

    // Deduplicate: if already fetching, don't start another request
    if (get()._isFetching) {
      return;
    }

    // Skip if we already have categories (unless force is true)
    if (!force && get().categories.length > 0) {
      return;
    }

    set({ isLoading: true, error: null, _isFetching: true });
    try {
      const response = await categoriesApi.getAll();
      if (response.categories) {
        set({ categories: response.categories });
      }
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Failed to fetch categories";
      set({ error: message });
    } finally {
      set({ isLoading: false, _isFetching: false });
    }
  },

  // Sync Actions
  setCategories: (categories) => set({ categories }),

  addCategory: (category) =>
    set((state) => ({
      categories: [...state.categories, category],
    })),

  updateCategory: (id, updates) =>
    set((state) => ({
      categories: state.categories.map((cat) =>
        cat.id === id ? { ...cat, ...updates } : cat,
      ),
    })),

  removeCategory: (id) =>
    set((state) => ({
      categories: state.categories.filter((cat) => cat.id !== id),
    })),

  clearCategories: () => set({ categories: [], _isFetching: false }),

  clearError: () => set({ error: null }),

  // Selectors
  getCategoryById: (id) => {
    const { categories } = get();
    return categories.find((cat) => cat.id === id);
  },

  getCategoriesByType: (type) => {
    const { categories } = get();
    return categories.filter((cat) => cat.type === type);
  },
}));

// -----------------------------------------------------------------------------
// Selector Hooks (for optimized re-renders)
// -----------------------------------------------------------------------------

export const useCategories = () =>
  useCategoriesStore((state) => state.categories);

export const useCategoriesLoading = () =>
  useCategoriesStore((state) => state.isLoading);

export const useCategoriesError = () =>
  useCategoriesStore((state) => state.error);

// Use shallow equality to prevent unnecessary re-renders when actions haven't changed
export const useCategoryActions = () =>
  useCategoriesStore(
    useShallow((state) => ({
      fetchCategories: state.fetchCategories,
      setCategories: state.setCategories,
      addCategory: state.addCategory,
      updateCategory: state.updateCategory,
      removeCategory: state.removeCategory,
      clearCategories: state.clearCategories,
      clearError: state.clearError,
    })),
  );
