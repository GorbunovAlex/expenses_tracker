"use client";

import { useEffect, useState } from "react";
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
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);
  const [isChecking, setIsChecking] = useState(true);

  useEffect(() => {
    // Small delay to allow hydration of persisted auth state
    const checkAuth = () => {
      if (isAuthenticated) {
        router.replace("/");
      } else {
        setIsChecking(false);
      }
    };

    // Use a small timeout to ensure Zustand has hydrated from localStorage
    const timeoutId = setTimeout(checkAuth, 50);

    return () => clearTimeout(timeoutId);
  }, [isAuthenticated, router]);

  // Show loading state while checking authentication
  if (isChecking) {
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

  // User is not authenticated, render children (guest content)
  return <>{children}</>;
}
