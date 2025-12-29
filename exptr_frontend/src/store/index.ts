// =============================================================================
// Store - Barrel Export
// =============================================================================

// Auth Store
export {
  useAuthStore,
  useUser,
  useIsAuthenticated,
} from "./authStore";

// Operations Store
export {
  useOperationsStore,
  useOperations,
  useOperationsActions,
} from "./operationsStore";

// Categories Store
export {
  useCategoriesStore,
  useCategories,
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
export {
  useUIStore,
  useIsLoading,
  useError,
  useSidebarOpen,
} from "./uiStore";
