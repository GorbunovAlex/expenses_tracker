import { create } from "zustand";
import { persist } from "zustand/middleware";

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

interface UIState {
  // State
  isLoading: boolean;
  error: string | null;
  sidebarOpen: boolean;

  // Actions
  setLoading: (isLoading: boolean) => void;
  setError: (error: string | null) => void;
  clearError: () => void;
  toggleSidebar: () => void;
  setSidebarOpen: (open: boolean) => void;
}

// -----------------------------------------------------------------------------
// Store
// -----------------------------------------------------------------------------

export const useUIStore = create<UIState>()(
  persist(
    (set) => ({
      // Initial State
      isLoading: false,
      error: null,
      sidebarOpen: true,

      // Actions
      setLoading: (isLoading) => set({ isLoading }),

      setError: (error) => set({ error }),

      clearError: () => set({ error: null }),

      toggleSidebar: () =>
        set((state) => ({
          sidebarOpen: !state.sidebarOpen,
        })),

      setSidebarOpen: (open) => set({ sidebarOpen: open }),
    }),
    {
      name: "exptr-ui-storage",
      partialize: (state) => ({
        sidebarOpen: state.sidebarOpen,
      }),
    },
  ),
);

// -----------------------------------------------------------------------------
// Selector Hooks
// -----------------------------------------------------------------------------

export const useIsLoading = () => useUIStore((state) => state.isLoading);
export const useError = () => useUIStore((state) => state.error);
export const useSidebarOpen = () => useUIStore((state) => state.sidebarOpen);
