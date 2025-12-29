// =============================================================================
// Base API Response Types
// =============================================================================

/**
 * Base API response structure returned by all endpoints
 */
export interface ApiResponse {
  status: string;
  error?: string;
}

/**
 * API Error class for handling HTTP errors
 */
export class ApiError extends Error {
  constructor(
    public status: number,
    message: string,
  ) {
    super(message);
    this.name = "ApiError";
  }
}
