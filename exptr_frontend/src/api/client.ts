/**
 * =============================================================================
 * API Client (Axios)
 * =============================================================================
 */

import axios, { AxiosError } from "axios";
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

/**
 * -----------------------------------------------------------------------------
 * Configuration
 * -----------------------------------------------------------------------------
 */

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "/api/v1";

/**
 * -----------------------------------------------------------------------------
 * Token Management
 * -----------------------------------------------------------------------------
 */

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

/**
 * -----------------------------------------------------------------------------
 * Axios Instance
 * -----------------------------------------------------------------------------
 */

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Request: attach Authorization header if token exists
api.interceptors.request.use((config) => {
  const token = getAuthToken();
  if (token) {
    config.headers.set("Authorization", `Bearer ${token}`);
  }
  return config;
});

// Response: handle 401 globally, normalize errors
api.interceptors.response.use(
  (response) => response,
  (error: AxiosError) => {
    const status = error.response?.status;

    if (status === 401) {
      removeAuthToken();
      authEvents.emitUnauthorized();
    }

    // Try to extract server-provided error message
    let message = "Request failed";
    if (error.response?.data) {
      if (typeof error.response.data === "string") {
        message = error.response.data;
      } else if (
        typeof error.response.data === "object" &&
        error.response.data !== null
      ) {
        const maybeMsg =
          (error.response.data as Record<string, unknown>).message ??
          (error.response.data as Record<string, unknown>).error;
        if (typeof maybeMsg === "string" && maybeMsg.trim().length > 0) {
          message = maybeMsg;
        }
      }
    } else if (error.message) {
      message = error.message;
    }

    throw new ApiError(status ?? 0, message);
  },
);

/**
 * -----------------------------------------------------------------------------
 * Base Request Helpers
 * -----------------------------------------------------------------------------
 */

enum HttpMethod {
  GET = "GET",
  POST = "POST",
  PUT = "PUT",
  DELETE = "DELETE",
}

async function request<T>(
  method: HttpMethod,
  endpoint: string,
  data?: unknown,
): Promise<T> {
  const res = await api.request<T>({
    method,
    url: endpoint,
    data,
  });
  return res.data;
}

/**
 * -----------------------------------------------------------------------------
 * Auth API
 * -----------------------------------------------------------------------------
 */

export const authApi = {
  login: async (data: LoginRequest): Promise<LoginResponse> => {
    const response = await request<LoginResponse>(
      HttpMethod.POST,
      "/users/login",
      data,
    );
    if (response.token) {
      setAuthToken(response.token);
    }
    return response;
  },

  signup: async (data: SignUpRequest): Promise<ApiResponse> => {
    return request<ApiResponse>(HttpMethod.POST, "/users/signup", data);
  },

  logout: (): void => {
    removeAuthToken();
  },

  isAuthenticated: (): boolean => {
    return !!getAuthToken();
  },
};

/**
 * -----------------------------------------------------------------------------
 * Categories API
 * -----------------------------------------------------------------------------
 */

export const categoriesApi = {
  getAll: async (): Promise<GetCategoriesResponse> => {
    return request<GetCategoriesResponse>(HttpMethod.GET, "/categories/");
  },

  create: async (data: CategoryRequest): Promise<CategoryResponse> => {
    return request<CategoryResponse>(HttpMethod.POST, "/categories/new", data);
  },

  update: async (id: string, data: CategoryRequest): Promise<ApiResponse> => {
    return request<ApiResponse>(HttpMethod.PUT, `/categories/${id}`, data);
  },

  delete: async (id: string): Promise<ApiResponse> => {
    return request<ApiResponse>(HttpMethod.DELETE, `/categories/${id}`);
  },
};

/**
 * -----------------------------------------------------------------------------
 * Operations API
 * -----------------------------------------------------------------------------
 */

export const operationsApi = {
  getAll: async (): Promise<GetOperationsResponse> => {
    return request<GetOperationsResponse>(HttpMethod.GET, "/operations");
  },

  create: async (data: OperationRequest): Promise<CreateOperationResponse> => {
    return request<CreateOperationResponse>(
      HttpMethod.POST,
      "/operations/new",
      data,
    );
  },

  update: async (
    id: string,
    data: Partial<OperationRequest>,
  ): Promise<UpdateOperationResponse> => {
    return request<UpdateOperationResponse>(
      HttpMethod.PUT,
      `/operations/${id}`,
      data,
    );
  },

  delete: async (id: string): Promise<ApiResponse> => {
    return request<ApiResponse>(HttpMethod.DELETE, `/operations/${id}`);
  },
};

/**
 * -----------------------------------------------------------------------------
 * Exports
 * -----------------------------------------------------------------------------
 */

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
