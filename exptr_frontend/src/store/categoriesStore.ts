import { create } from "zustand";
import type { Category } from "@/types/api";

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

interface CategoriesState {
  // State
  categories: Category[];

  // Actions
  setCategories: (categories: Category[]) => void;
  addCategory: (category: Category) => void;
  updateCategory: (id: string, updates: Partial<Category>) => void;
  removeCategory: (id: string) => void;
  clearCategories: () => void;

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

  // Actions
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

  clearCategories: () => set({ categories: [] }),

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

export const useCategoryActions = () =>
  useCategoriesStore((state) => ({
    setCategories: state.setCategories,
    addCategory: state.addCategory,
    updateCategory: state.updateCategory,
    removeCategory: state.removeCategory,
    clearCategories: state.clearCategories,
  }));
