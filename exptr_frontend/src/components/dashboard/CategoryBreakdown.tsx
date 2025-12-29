"use client";

import * as React from "react";
import {
  PieChart,
  Pie,
  Cell,
  ResponsiveContainer,
  Legend,
  Tooltip,
} from "recharts";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import type { Category, Operation } from "@/types/api";

interface CategoryBreakdownProps {
  operations: Operation[];
  categories: Category[];
  type?: "expense" | "income";
}

interface CategoryData {
  name: string;
  value: number;
  color: string;
  icon: string;
  percentage: number;
}

const formatCurrencyValue = (value: number) => {
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(value / 100);
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const CategoryTooltip = ({ active, payload }: any) => {
  if (active && payload && payload.length) {
    const data = payload[0].payload as CategoryData;
    return (
      <div className="rounded-md border border-border bg-popover px-3 py-2 shadow-lg">
        <p className="font-medium text-foreground">
          {data.icon} {data.name}
        </p>
        <p className="text-sm text-muted-foreground">
          {formatCurrencyValue(data.value)} ({data.percentage.toFixed(1)}%)
        </p>
      </div>
    );
  }
  return null;
};

const DEFAULT_COLORS = [
  "#0FB5B3", // Glacial Teal
  "#006D80", // Deep Cerulean
  "#38bdf8", // Sky
  "#2dd4bf", // Teal
  "#a78bfa", // Violet
  "#f472b6", // Pink
  "#fb923c", // Orange
  "#facc15", // Yellow
];

export function CategoryBreakdown({
  operations,
  categories,
  type = "expense",
}: CategoryBreakdownProps) {
  const categoryData = React.useMemo(() => {
    // Filter operations by type
    const filteredOps = operations.filter((op) => op.type === type);

    // Group by category
    const categoryTotals = new Map<string, number>();
    filteredOps.forEach((op) => {
      const current = categoryTotals.get(op.category_id) || 0;
      categoryTotals.set(op.category_id, current + op.amount);
    });

    // Calculate total for percentages
    const total = Array.from(categoryTotals.values()).reduce(
      (sum, val) => sum + val,
      0,
    );

    // Build category data with percentages
    const data: CategoryData[] = [];
    categoryTotals.forEach((value, categoryId) => {
      const category = categories.find((c) => c.id === categoryId);
      if (category) {
        data.push({
          name: category.name,
          value,
          color:
            category.color ||
            DEFAULT_COLORS[data.length % DEFAULT_COLORS.length],
          icon: category.icon || "📦",
          percentage: total > 0 ? (value / total) * 100 : 0,
        });
      }
    });

    // Sort by value descending
    return data.sort((a, b) => b.value - a.value);
  }, [operations, categories, type]);

  const total = categoryData.reduce((sum, item) => sum + item.value, 0);

  const formatCurrency = (value: number) => formatCurrencyValue(value);

  if (categoryData.length === 0) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="text-lg">
            {type === "expense" ? "Expense" : "Income"} by Category
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex h-50 items-center justify-center text-muted-foreground">
            No {type} data available
          </div>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-lg">
          {type === "expense" ? "Expense" : "Income"} by Category
        </CardTitle>
      </CardHeader>
      <CardContent>
        <div className="grid gap-6 md:grid-cols-2">
          {/* Pie Chart */}
          <div className="h-64">
            <ResponsiveContainer width="100%" height="100%">
              <PieChart>
                <Pie
                  data={
                    categoryData as unknown as Array<Record<string, unknown>>
                  }
                  cx="50%"
                  cy="50%"
                  innerRadius={60}
                  outerRadius={90}
                  paddingAngle={2}
                  dataKey="value"
                >
                  {categoryData.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.color} />
                  ))}
                </Pie>
                <Tooltip content={<CategoryTooltip />} />
                <Legend
                  formatter={(value) => (
                    <span className="text-sm text-foreground">{value}</span>
                  )}
                />
              </PieChart>
            </ResponsiveContainer>
          </div>

          {/* Category List with Progress */}
          <div className="space-y-4">
            <div className="text-sm font-medium text-muted-foreground">
              Total: {formatCurrency(total)}
            </div>
            <div className="space-y-3">
              {categoryData.slice(0, 5).map((item) => (
                <div key={item.name} className="space-y-1">
                  <div className="flex items-center justify-between text-sm">
                    <span className="flex items-center gap-2">
                      <span>{item.icon}</span>
                      <span className="font-medium">{item.name}</span>
                    </span>
                    <span className="text-muted-foreground">
                      {formatCurrency(item.value)}
                    </span>
                  </div>
                  <Progress
                    value={item.percentage}
                    className="h-2"
                    indicatorClassName={`bg-[${item.color}]`}
                    style={
                      {
                        "--progress-color": item.color,
                      } as React.CSSProperties
                    }
                  />
                </div>
              ))}
              {categoryData.length > 5 && (
                <p className="text-xs text-muted-foreground">
                  +{categoryData.length - 5} more categories
                </p>
              )}
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
