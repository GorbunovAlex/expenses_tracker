// =============================================================================
// User & Authentication Types
// =============================================================================

import type { ApiResponse } from "./common";

// -----------------------------------------------------------------------------
// Request Types
// -----------------------------------------------------------------------------

export interface LoginRequest {
  email: string;
  password: string;
}

export interface SignUpRequest {
  email: string;
  password: string;
}

// -----------------------------------------------------------------------------
// Response Types
// -----------------------------------------------------------------------------

export interface LoginResponse extends ApiResponse {
  token?: string;
}

export type SignUpResponse = ApiResponse;

// -----------------------------------------------------------------------------
// User Entity
// -----------------------------------------------------------------------------

export interface User {
  id: string;
  email: string;
  created_at: string;
  updated_at: string;
}

// -----------------------------------------------------------------------------
// Auth State Types (for frontend store)
// -----------------------------------------------------------------------------

export interface AuthUser {
  email: string;
  token: string;
}
