// =============================================================================
// API Client
// =============================================================================

import { authEvents } from "@/lib/authEvents";
import {
  ApiError,
  type ApiResponse,
  type LoginRequest,
  type LoginResponse,
  type SignUpRequest,
  type Category,
  type CategoryRequest,
  type CategoryResponse,
  type GetCategoriesResponse,
  type Operation,
  type OperationRequest,
  type CreateOperationResponse,
  type UpdateOperationResponse,
  type GetOperationsResponse,
} from "@/types/api";

// -----------------------------------------------------------------------------
// Configuration
// -----------------------------------------------------------------------------

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "/api/v1";

// -----------------------------------------------------------------------------
// Token Management
// -----------------------------------------------------------------------------

function getAuthToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("auth_token");
}

function setAuthToken(token: string): void {
  if (typeof window !== "undefined") {
    localStorage.setItem("auth_token", token);
  }
}

function removeAuthToken(): void {
  if (typeof window !== "undefined") {
    localStorage.removeItem("auth_token");
  }
}

// -----------------------------------------------------------------------------
// Base Request Function
// -----------------------------------------------------------------------------

async function request<T>(
  endpoint: string,
  options: RequestInit = {},
): Promise<T> {
  const token = getAuthToken();
  const headers: HeadersInit = {
    "Content-Type": "application/json",
    ...options.headers,
  };

  if (token) {
    (headers as Record<string, string>)["Authorization"] = `Bearer ${token}`;
  }

  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    // Handle 401 Unauthorized globally
    if (response.status === 401) {
      removeAuthToken();
      authEvents.emitUnauthorized();
    }

    const errorText = await response.text();
    throw new ApiError(response.status, errorText || "Request failed");
  }

  return response.json();
}

// -----------------------------------------------------------------------------
// Auth API
// -----------------------------------------------------------------------------

export const authApi = {
  login: async (data: LoginRequest): Promise<LoginResponse> => {
    const response = await request<LoginResponse>("/users/login", {
      method: "POST",
      body: JSON.stringify(data),
    });
    if (response.token) {
      setAuthToken(response.token);
    }
    return response;
  },

  signup: async (data: SignUpRequest): Promise<ApiResponse> => {
    return request<ApiResponse>("/users/signup", {
      method: "POST",
      body: JSON.stringify(data),
    });
  },

  logout: (): void => {
    removeAuthToken();
  },

  isAuthenticated: (): boolean => {
    return !!getAuthToken();
  },
};

// -----------------------------------------------------------------------------
// Categories API
// -----------------------------------------------------------------------------

export const categoriesApi = {
  getAll: async (): Promise<GetCategoriesResponse> => {
    return request<GetCategoriesResponse>("/categories/");
  },

  create: async (data: CategoryRequest): Promise<CategoryResponse> => {
    return request<CategoryResponse>("/categories/new", {
      method: "POST",
      body: JSON.stringify(data),
    });
  },

  update: async (id: string, data: CategoryRequest): Promise<ApiResponse> => {
    return request<ApiResponse>(`/categories/${id}`, {
      method: "PUT",
      body: JSON.stringify(data),
    });
  },

  delete: async (id: string): Promise<ApiResponse> => {
    return request<ApiResponse>(`/categories/${id}`, {
      method: "DELETE",
    });
  },
};

// -----------------------------------------------------------------------------
// Operations API
// -----------------------------------------------------------------------------

export const operationsApi = {
  getAll: async (): Promise<GetOperationsResponse> => {
    return request<GetOperationsResponse>("/operations");
  },

  create: async (data: OperationRequest): Promise<CreateOperationResponse> => {
    return request<CreateOperationResponse>("/operations/new", {
      method: "POST",
      body: JSON.stringify(data),
    });
  },

  update: async (
    id: string,
    data: Partial<OperationRequest>,
  ): Promise<UpdateOperationResponse> => {
    return request<UpdateOperationResponse>(`/operations/${id}`, {
      method: "PUT",
      body: JSON.stringify(data),
    });
  },

  delete: async (id: string): Promise<ApiResponse> => {
    return request<ApiResponse>(`/operations/${id}`, {
      method: "DELETE",
    });
  },
};

// -----------------------------------------------------------------------------
// Exports
// -----------------------------------------------------------------------------

// Re-export types for backward compatibility
export type {
  ApiResponse,
  LoginRequest,
  LoginResponse,
  SignUpRequest,
  Category,
  CategoryRequest,
  CategoryResponse,
  GetCategoriesResponse,
  Operation,
  OperationRequest,
  CreateOperationResponse,
  UpdateOperationResponse,
  GetOperationsResponse,
};

export { ApiError, getAuthToken, setAuthToken, removeAuthToken };
