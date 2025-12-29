import { create } from "zustand";
import { persist } from "zustand/middleware";

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

export type DateRange = "week" | "month" | "year" | "all";
export type FilterType = "all" | "income" | "expense";

export interface DashboardFilters {
  dateRange: DateRange;
  categoryId: string | null;
  type: FilterType;
}

interface FiltersState {
  // State
  filters: DashboardFilters;

  // Actions
  setFilters: (filters: Partial<DashboardFilters>) => void;
  setDateRange: (dateRange: DateRange) => void;
  setCategoryId: (categoryId: string | null) => void;
  setType: (type: FilterType) => void;
  resetFilters: () => void;
}

// -----------------------------------------------------------------------------
// Default Values
// -----------------------------------------------------------------------------

const defaultFilters: DashboardFilters = {
  dateRange: "month",
  categoryId: null,
  type: "all",
};

// -----------------------------------------------------------------------------
// Store
// -----------------------------------------------------------------------------

export const useFiltersStore = create<FiltersState>()(
  persist(
    (set) => ({
      // Initial State
      filters: defaultFilters,

      // Actions
      setFilters: (filters) =>
        set((state) => ({
          filters: { ...state.filters, ...filters },
        })),

      setDateRange: (dateRange) =>
        set((state) => ({
          filters: { ...state.filters, dateRange },
        })),

      setCategoryId: (categoryId) =>
        set((state) => ({
          filters: { ...state.filters, categoryId },
        })),

      setType: (type) =>
        set((state) => ({
          filters: { ...state.filters, type },
        })),

      resetFilters: () => set({ filters: defaultFilters }),
    }),
    {
      name: "exptr-filters-storage",
    },
  ),
);

// -----------------------------------------------------------------------------
// Selector Hooks
// -----------------------------------------------------------------------------

export const useFilters = () => useFiltersStore((state) => state.filters);
export const useDateRange = () =>
  useFiltersStore((state) => state.filters.dateRange);
export const useFilterCategoryId = () =>
  useFiltersStore((state) => state.filters.categoryId);
export const useFilterType = () =>
  useFiltersStore((state) => state.filters.type);
