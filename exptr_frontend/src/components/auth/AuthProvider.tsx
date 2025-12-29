"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { authEvents } from "@/lib/authEvents";
import { useAuthStore } from "@/store/authStore";

interface AuthProviderProps {
  children: React.ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const router = useRouter();
  const logout = useAuthStore((state) => state.logout);

  useEffect(() => {
    // Subscribe to unauthorized events from the API client
    const unsubscribe = authEvents.onUnauthorized(() => {
      // Clear auth state
      logout();

      // Redirect to login with a message
      router.replace("/login?session=expired");
    });

    return () => {
      unsubscribe();
    };
  }, [router, logout]);

  return <>{children}</>;
}
