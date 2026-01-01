// =============================================================================
// Store - Barrel Export
// =============================================================================

// Auth Store
export {
  useAuthStore,
  useUser,
  useIsAuthenticated,
  useHasHydrated,
} from "./authStore";

// Operations Store
export {
  useOperationsStore,
  useOperations,
  useOperationsLoading,
  useOperationsError,
  useOperationsActions,
} from "./operationsStore";

// Categories Store
export {
  useCategoriesStore,
  useCategories,
  useCategoriesLoading,
  useCategoriesError,
  useCategoryActions,
} from "./categoriesStore";

// Filters Store
export {
  useFiltersStore,
  useFilters,
  useDateRange,
  useFilterCategoryId,
  useFilterType,
  type DashboardFilters,
  type DateRange,
  type FilterType,
} from "./filtersStore";

// UI Store
export { useUIStore, useIsLoading, useError, useSidebarOpen } from "./uiStore";
