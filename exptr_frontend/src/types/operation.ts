// =============================================================================
// Operation Types
// =============================================================================

import type { ApiResponse } from "./common";

// -----------------------------------------------------------------------------
// Enums / Literal Types
// -----------------------------------------------------------------------------

export type OperationType = "income" | "expense";
export type Currency = "USD" | "EUR" | "GBP" | "RUB" | string;

// -----------------------------------------------------------------------------
// Entity Types
// -----------------------------------------------------------------------------

export interface Operation {
  id: string;
  name: string;
  type: OperationType;
  amount: number;
  currency: Currency;
  category_id: string;
  comment?: string;
  user_id: string;
  created_at: string;
  updated_at: string;
}

// -----------------------------------------------------------------------------
// Request Types
// -----------------------------------------------------------------------------

export interface OperationRequest {
  name: string;
  type: OperationType;
  amount: number;
  currency: Currency;
  category_id: string;
  comment?: string;
  user_id: string;
  created_at: string;
  updated_at?: string;
}

// -----------------------------------------------------------------------------
// Response Types
// -----------------------------------------------------------------------------

export type CreateOperationResponse = ApiResponse;

export type UpdateOperationResponse = ApiResponse;

export interface GetOperationsResponse extends ApiResponse {
  operations?: Operation[];
}
