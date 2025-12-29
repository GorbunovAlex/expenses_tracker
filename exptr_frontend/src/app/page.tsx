"use client";

import * as React from "react";
import {
  TrendingUp,
  TrendingDown,
  Wallet,
  ArrowUpDown,
  Plus,
} from "lucide-react";

import { AuthGuard } from "@/components/auth";
import { useAuthStore } from "@/store/authStore";
import { useOperationsStore } from "@/store/operationsStore";
import { useCategoriesStore } from "@/store/categoriesStore";
import { useUIStore } from "@/store/uiStore";
import { useFiltersStore } from "@/store/filtersStore";

import {
  operationsApi,
  categoriesApi,
  type Operation,
  type OperationRequest,
} from "@/api/client";

import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

import { Sidebar } from "@/components/dashboard/Sidebar";
import { Header } from "@/components/dashboard/Header";
import { StatCard } from "@/components/dashboard/StatCard";
import { RecentTransactions } from "@/components/dashboard/RecentTransactions";
import { CategoryBreakdown } from "@/components/dashboard/CategoryBreakdown";
import { MonthlyChart } from "@/components/dashboard/MonthlyChart";
import { AddOperationDialog } from "@/components/dashboard/AddOperationDialog";

function DashboardContent() {
  // Auth store
  const user = useAuthStore((state) => state.user);
  const logout = useAuthStore((state) => state.logout);

  // Operations store
  const operations = useOperationsStore((state) => state.operations);
  const setOperations = useOperationsStore((state) => state.setOperations);
  const removeOperation = useOperationsStore((state) => state.removeOperation);
  const clearOperations = useOperationsStore((state) => state.clearOperations);

  // Categories store
  const categories = useCategoriesStore((state) => state.categories);
  const setCategories = useCategoriesStore((state) => state.setCategories);
  const clearCategories = useCategoriesStore((state) => state.clearCategories);

  // UI store
  const isLoading = useUIStore((state) => state.isLoading);
  const setLoading = useUIStore((state) => state.setLoading);
  const error = useUIStore((state) => state.error);
  const setError = useUIStore((state) => state.setError);
  const sidebarOpen = useUIStore((state) => state.sidebarOpen);
  const toggleSidebar = useUIStore((state) => state.toggleSidebar);

  // Filters store
  const resetFilters = useFiltersStore((state) => state.resetFilters);

  const [isAddDialogOpen, setIsAddDialogOpen] = React.useState(false);
  const [selectedOperation, setSelectedOperation] =
    React.useState<Operation | null>(null);

  // Handle logout with cleanup
  const handleLogout = React.useCallback(() => {
    logout();
    clearOperations();
    clearCategories();
    resetFilters();
  }, [logout, clearOperations, clearCategories, resetFilters]);

  // Fetch data on mount
  React.useEffect(() => {
    async function fetchData() {
      setLoading(true);
      setError(null);

      try {
        const [opsResponse, catsResponse] = await Promise.all([
          operationsApi.getAll(),
          categoriesApi.getAll(),
        ]);

        if (opsResponse.operations) {
          setOperations(opsResponse.operations);
        }

        if (catsResponse.categories) {
          setCategories(catsResponse.categories);
        }
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to fetch data");
      } finally {
        setLoading(false);
      }
    }

    fetchData();
  }, [setOperations, setCategories, setLoading, setError]);

  // Calculate statistics
  const stats = React.useMemo(() => {
    const totalIncome = operations
      .filter((op) => op.type === "income")
      .reduce((sum, op) => sum + op.amount, 0);

    const totalExpense = operations
      .filter((op) => op.type === "expense")
      .reduce((sum, op) => sum + op.amount, 0);

    const balance = totalIncome - totalExpense;

    return {
      totalIncome: totalIncome / 100,
      totalExpense: totalExpense / 100,
      balance: balance / 100,
      transactionCount: operations.length,
    };
  }, [operations]);

  const formatCurrency = (value: number): string => {
    return new Intl.NumberFormat("en-US", {
      style: "currency",
      currency: "USD",
      minimumFractionDigits: 0,
      maximumFractionDigits: 2,
    }).format(value);
  };

  const handleAddOperation = async (data: OperationRequest) => {
    const response = await operationsApi.create(data);
    if (response.status === "OK" || response.status === "success") {
      // Refetch operations to get the new one with ID
      const opsResponse = await operationsApi.getAll();
      if (opsResponse.operations) {
        setOperations(opsResponse.operations);
      }
    }
  };

  const handleDeleteOperation = async (operation: Operation) => {
    if (!confirm("Are you sure you want to delete this transaction?")) {
      return;
    }

    try {
      await operationsApi.delete(operation.id);
      removeOperation(operation.id);
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Failed to delete transaction",
      );
    }
  };

  const handleEditOperation = (operation: Operation) => {
    setSelectedOperation(operation);
    // TODO: Open edit dialog
    console.log("Edit operation:", operation);
  };

  return (
    <div className="flex min-h-screen bg-background">
      {/* Sidebar */}
      <Sidebar
        isCollapsed={!sidebarOpen}
        onToggle={toggleSidebar}
        userEmail={user?.email}
        onLogout={handleLogout}
      />

      {/* Main Content */}
      <div className="flex flex-1 flex-col overflow-hidden">
        {/* Header */}
        <Header
          onMenuClick={toggleSidebar}
          onAddTransaction={() => setIsAddDialogOpen(true)}
          userEmail={user?.email}
          notificationCount={2}
        />

        {/* Page Content */}
        <main className="flex-1 overflow-y-auto p-4 md:p-6 lg:p-8">
          {/* Page Title */}
          <div className="mb-8 flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
            <div>
              <h1 className="text-2xl font-bold tracking-tight md:text-3xl">
                Dashboard
              </h1>
              <p className="text-muted-foreground">
                Welcome back! Here&apos;s your financial overview.
              </p>
            </div>
            <Button
              onClick={() => setIsAddDialogOpen(true)}
              className="w-full gap-2 md:w-auto"
            >
              <Plus className="h-4 w-4" />
              Add Transaction
            </Button>
          </div>

          {/* Error State */}
          {error && (
            <div className="mb-6 rounded-lg border border-destructive/50 bg-destructive/10 px-4 py-3 text-sm text-destructive">
              {error}
            </div>
          )}

          {/* Loading State */}
          {isLoading ? (
            <div className="flex h-96 items-center justify-center">
              <div className="text-center">
                <div className="mx-auto mb-4 h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent" />
                <p className="text-muted-foreground">Loading your data...</p>
              </div>
            </div>
          ) : (
            <>
              {/* Stats Cards */}
              <div className="mb-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
                <StatCard
                  title="Total Balance"
                  value={formatCurrency(stats.balance)}
                  subtitle="Current balance"
                  icon={Wallet}
                  variant="balance"
                  trend={
                    stats.balance > 0
                      ? { value: 12.5, isPositive: true }
                      : undefined
                  }
                />
                <StatCard
                  title="Total Income"
                  value={formatCurrency(stats.totalIncome)}
                  subtitle="This month"
                  icon={TrendingUp}
                  variant="income"
                />
                <StatCard
                  title="Total Expenses"
                  value={formatCurrency(stats.totalExpense)}
                  subtitle="This month"
                  icon={TrendingDown}
                  variant="expense"
                />
                <StatCard
                  title="Transactions"
                  value={stats.transactionCount.toString()}
                  subtitle="Total operations"
                  icon={ArrowUpDown}
                  variant="default"
                />
              </div>

              {/* Charts Section */}
              <div className="mb-8 grid gap-6 lg:grid-cols-2">
                <MonthlyChart operations={operations} />

                <Tabs defaultValue="expense" className="w-full">
                  <TabsList className="mb-4">
                    <TabsTrigger value="expense">Expenses</TabsTrigger>
                    <TabsTrigger value="income">Income</TabsTrigger>
                  </TabsList>
                  <TabsContent value="expense">
                    <CategoryBreakdown
                      operations={operations}
                      categories={categories}
                      type="expense"
                    />
                  </TabsContent>
                  <TabsContent value="income">
                    <CategoryBreakdown
                      operations={operations}
                      categories={categories}
                      type="income"
                    />
                  </TabsContent>
                </Tabs>
              </div>

              {/* Recent Transactions */}
              <RecentTransactions
                operations={operations}
                categories={categories}
                limit={10}
                onEdit={handleEditOperation}
                onDelete={handleDeleteOperation}
              />
            </>
          )}
        </main>
      </div>

      {/* Add Operation Dialog */}
      <AddOperationDialog
        open={isAddDialogOpen}
        onOpenChange={setIsAddDialogOpen}
        categories={categories}
        onSubmit={handleAddOperation}
        userId={user?.email || "anonymous"}
      />
    </div>
  );
}

export default function DashboardPage() {
  return (
    <AuthGuard>
      <DashboardContent />
    </AuthGuard>
  );
}
