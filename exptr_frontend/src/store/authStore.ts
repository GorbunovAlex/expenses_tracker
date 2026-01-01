import { create } from "zustand";
import { persist } from "zustand/middleware";
import type { AuthUser } from "@/types/api";

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

interface AuthState {
  user: AuthUser | null;
  isAuthenticated: boolean;
  _hasHydrated: boolean;

  // Actions
  setUser: (user: AuthUser | null) => void;
  login: (email: string, token: string) => void;
  logout: () => void;
  setHasHydrated: (hasHydrated: boolean) => void;
}

// -----------------------------------------------------------------------------
// Store
// -----------------------------------------------------------------------------

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      // Initial State
      user: null,
      isAuthenticated: false,
      _hasHydrated: false,

      // Actions
      setUser: (user) =>
        set({
          user,
          isAuthenticated: !!user,
        }),

      login: (email, token) =>
        set({
          user: { email, token },
          isAuthenticated: true,
        }),

      logout: () => {
        if (typeof window !== "undefined") {
          localStorage.removeItem("auth_token");
        }
        set({
          user: null,
          isAuthenticated: false,
        });
      },

      setHasHydrated: (hasHydrated) => set({ _hasHydrated: hasHydrated }),
    }),
    {
      name: "exptr-auth-storage",
      partialize: (state) => ({
        user: state.user,
        isAuthenticated: state.isAuthenticated,
      }),
      onRehydrateStorage: () => (state) => {
        // Called when hydration is complete
        state?.setHasHydrated(true);
      },
    },
  ),
);

// -----------------------------------------------------------------------------
// Selector Hooks
// -----------------------------------------------------------------------------

export const useUser = () => useAuthStore((state) => state.user);
export const useIsAuthenticated = () =>
  useAuthStore((state) => state.isAuthenticated);
export const useHasHydrated = () => useAuthStore((state) => state._hasHydrated);
