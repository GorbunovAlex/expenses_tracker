// =============================================================================
// Dashboard Statistics Types
// =============================================================================

import type { Category } from "./category";

// -----------------------------------------------------------------------------
// Dashboard Stats
// -----------------------------------------------------------------------------

export interface DashboardStats {
  totalIncome: number;
  totalExpense: number;
  balance: number;
  transactionCount: number;
}

// -----------------------------------------------------------------------------
// Category Statistics
// -----------------------------------------------------------------------------

export interface CategoryStats {
  category: Category;
  total: number;
  count: number;
  percentage: number;
}

// -----------------------------------------------------------------------------
// Monthly Statistics
// -----------------------------------------------------------------------------

export interface MonthlyStats {
  month: string;
  income: number;
  expense: number;
}
