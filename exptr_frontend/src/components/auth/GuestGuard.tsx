"use client";

import { useEffect, useRef } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/authStore";

interface GuestGuardProps {
  children: React.ReactNode;
  fallback?: React.ReactNode;
}

/**
 * GuestGuard - Protects guest-only routes (login, signup)
 * Redirects authenticated users to the home page
 */
export function GuestGuard({ children, fallback }: GuestGuardProps) {
  const router = useRouter();
  const hasRedirected = useRef(false);

  // Select state individually to avoid unnecessary re-renders
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  const hasHydrated = useAuthStore((state) => state._hasHydrated);

  useEffect(() => {
    // Wait for hydration to complete before making any routing decisions
    if (!hasHydrated) {
      return;
    }

    // Prevent duplicate redirects
    if (hasRedirected.current) {
      return;
    }

    // If authenticated after hydration, redirect to home
    if (isAuthenticated) {
      hasRedirected.current = true;
      router.replace("/");
    }
  }, [hasHydrated, isAuthenticated, router]);

  // Reset redirect flag if user becomes unauthenticated (e.g., after logout)
  useEffect(() => {
    if (!isAuthenticated) {
      hasRedirected.current = false;
    }
  }, [isAuthenticated]);

  // Show loading while waiting for hydration
  if (!hasHydrated) {
    return (
      fallback || (
        <div className="flex min-h-screen items-center justify-center bg-background">
          <div className="text-center">
            <div className="mx-auto mb-4 h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent" />
            <p className="text-muted-foreground">Loading...</p>
          </div>
        </div>
      )
    );
  }

  // Show loading while redirecting authenticated user
  if (isAuthenticated) {
    return (
      fallback || (
        <div className="flex min-h-screen items-center justify-center bg-background">
          <div className="text-center">
            <div className="mx-auto mb-4 h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent" />
            <p className="text-muted-foreground">Redirecting...</p>
          </div>
        </div>
      )
    );
  }

  // User is not authenticated, render children (guest content)
  return <>{children}</>;
}
