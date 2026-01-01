"use client";

import { useEffect, useRef } from "react";
import { useRouter, usePathname } from "next/navigation";
import { authEvents } from "@/lib/authEvents";
import { useAuthStore } from "@/store/authStore";

interface AuthProviderProps {
  children: React.ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const router = useRouter();
  const pathname = usePathname();
  const logout = useAuthStore((state) => state.logout);
  const isRedirecting = useRef(false);

  useEffect(() => {
    // Subscribe to unauthorized events from the API client
    const unsubscribe = authEvents.onUnauthorized(() => {
      // Prevent duplicate redirects
      if (isRedirecting.current) {
        return;
      }

      // Don't redirect if already on login page
      if (pathname === "/login" || pathname === "/signup") {
        return;
      }

      isRedirecting.current = true;

      // Clear auth state
      logout();

      // Redirect to login with a message
      router.replace("/login?session=expired");

      // Reset flag after navigation completes
      setTimeout(() => {
        isRedirecting.current = false;
      }, 1000);
    });

    return () => {
      unsubscribe();
    };
  }, [router, logout, pathname]);

  return <>{children}</>;
}
