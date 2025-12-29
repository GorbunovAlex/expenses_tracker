// =============================================================================
// Category Types
// =============================================================================

import type { ApiResponse } from "./common";

// -----------------------------------------------------------------------------
// Category Entity
// -----------------------------------------------------------------------------

export type CategoryType = "income" | "expense";

export interface Category {
  id: string;
  name: string;
  type: CategoryType;
  icon: string;
  color: string;
  user_id: string;
  created_at: string;
  updated_at: string;
}

// -----------------------------------------------------------------------------
// Category Request/Response Types
// -----------------------------------------------------------------------------

export interface CategoryRequest {
  name: string;
  type: CategoryType;
  icon?: string;
  color?: string;
  user_id?: string;
  created_at?: string;
  updated_at?: string;
}

export type CategoryResponse = ApiResponse;

export interface GetCategoriesResponse extends ApiResponse {
  categories?: Category[];
}
