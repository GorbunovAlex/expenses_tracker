// =============================================================================
// API Types - Barrel Export
// =============================================================================

// Common/Base types
export { ApiError, type ApiResponse } from "./common";

// User & Authentication types
export type {
  LoginRequest,
  LoginResponse,
  SignUpRequest,
  SignUpResponse,
  User,
  AuthUser,
} from "./user";

// Category types
export type {
  Category,
  CategoryType,
  CategoryRequest,
  CategoryResponse,
  GetCategoriesResponse,
} from "./category";

// Operation types
export type {
  Operation,
  OperationType,
  Currency,
  OperationRequest,
  CreateOperationResponse,
  UpdateOperationResponse,
  GetOperationsResponse,
} from "./operation";

// Dashboard statistics types
export type { DashboardStats, MonthlyStats, CategoryStats } from "./dashboard";

// -----------------------------------------------------------------------------
// Enriched Types (with joined data)
// -----------------------------------------------------------------------------

import type { Operation } from "./operation";
import type { Category } from "./category";

export interface OperationWithCategory extends Operation {
  category?: Category;
}
