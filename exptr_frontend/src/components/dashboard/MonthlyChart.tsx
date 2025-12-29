"use client";

import * as React from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  Legend,
} from "recharts";
import { format, parseISO, startOfMonth, subMonths } from "date-fns";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

const formatCurrencyValue = (value: number): string => {
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(value);
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const MonthlyTooltip = ({ active, payload, label }: any) => {
  if (active && payload && payload.length) {
    const incomeItem = payload.find(
      (p: { dataKey?: string }) => p.dataKey === "income",
    );
    const expenseItem = payload.find(
      (p: { dataKey?: string }) => p.dataKey === "expense",
    );
    const income = typeof incomeItem?.value === "number" ? incomeItem.value : 0;
    const expense =
      typeof expenseItem?.value === "number" ? expenseItem.value : 0;
    const net = income - expense;

    return (
      <div className="rounded-lg border border-border bg-popover p-3 shadow-lg">
        <p className="mb-2 font-medium text-foreground">{label}</p>
        <div className="space-y-1 text-sm">
          <div className="flex items-center justify-between gap-4">
            <span className="flex items-center gap-2">
              <span className="h-2 w-2 rounded-full bg-emerald-500" />
              Income
            </span>
            <span className="font-medium text-emerald-600">
              {formatCurrencyValue(income)}
            </span>
          </div>
          <div className="flex items-center justify-between gap-4">
            <span className="flex items-center gap-2">
              <span className="h-2 w-2 rounded-full bg-rose-500" />
              Expense
            </span>
            <span className="font-medium text-rose-600">
              {formatCurrencyValue(expense)}
            </span>
          </div>
          <div className="border-t border-border pt-1 mt-1">
            <div className="flex items-center justify-between gap-4">
              <span className="text-muted-foreground">Net</span>
              <span
                className={`font-semibold ${
                  net >= 0 ? "text-emerald-600" : "text-rose-600"
                }`}
              >
                {formatCurrencyValue(net)}
              </span>
            </div>
          </div>
        </div>
      </div>
    );
  }
  return null;
};
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import type { Operation } from "@/types/api";

interface MonthlyChartProps {
  operations: Operation[];
  className?: string;
}

interface MonthlyData {
  month: string;
  monthLabel: string;
  income: number;
  expense: number;
}

type TimeRange = "6" | "12";

export function MonthlyChart({ operations, className }: MonthlyChartProps) {
  const [timeRange, setTimeRange] = React.useState<TimeRange>("6");

  const monthlyData = React.useMemo(() => {
    const months = parseInt(timeRange);
    const now = new Date();
    const data: MonthlyData[] = [];

    // Initialize months
    for (let i = months - 1; i >= 0; i--) {
      const monthDate = startOfMonth(subMonths(now, i));
      const monthKey = format(monthDate, "yyyy-MM");
      data.push({
        month: monthKey,
        monthLabel: format(monthDate, "MMM yyyy"),
        income: 0,
        expense: 0,
      });
    }

    // Aggregate operations by month
    operations.forEach((op) => {
      try {
        const opDate = parseISO(op.created_at);
        const monthKey = format(opDate, "yyyy-MM");
        const monthData = data.find((d) => d.month === monthKey);

        if (monthData) {
          if (op.type === "income") {
            monthData.income += op.amount / 100; // Convert from cents
          } else if (op.type === "expense") {
            monthData.expense += op.amount / 100;
          }
        }
      } catch {
        // Skip invalid dates
      }
    });

    return data;
  }, [operations, timeRange]);

  const formatCurrency = (value: number): string => formatCurrencyValue(value);

  const totals = React.useMemo(() => {
    return monthlyData.reduce(
      (acc, month) => ({
        income: acc.income + month.income,
        expense: acc.expense + month.expense,
      }),
      { income: 0, expense: 0 },
    );
  }, [monthlyData]);

  return (
    <Card className={className}>
      <CardHeader className="flex flex-row items-center justify-between pb-2">
        <div>
          <CardTitle className="text-lg">Monthly Overview</CardTitle>
          <p className="text-sm text-muted-foreground mt-1">
            Income vs Expenses over the last {timeRange} months
          </p>
        </div>
        <Select
          value={timeRange}
          onValueChange={(v) => setTimeRange(v as TimeRange)}
        >
          <SelectTrigger className="w-[140px]">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="6">Last 6 months</SelectItem>
            <SelectItem value="12">Last 12 months</SelectItem>
          </SelectContent>
        </Select>
      </CardHeader>
      <CardContent>
        {/* Summary Stats */}
        <div className="mb-6 grid grid-cols-3 gap-4">
          <div className="rounded-lg bg-emerald-50 dark:bg-emerald-900/20 p-3">
            <p className="text-xs font-medium text-emerald-600 dark:text-emerald-400">
              Total Income
            </p>
            <p className="text-lg font-bold text-emerald-700 dark:text-emerald-300">
              {formatCurrency(totals.income)}
            </p>
          </div>
          <div className="rounded-lg bg-rose-50 dark:bg-rose-900/20 p-3">
            <p className="text-xs font-medium text-rose-600 dark:text-rose-400">
              Total Expenses
            </p>
            <p className="text-lg font-bold text-rose-700 dark:text-rose-300">
              {formatCurrency(totals.expense)}
            </p>
          </div>
          <div className="rounded-lg bg-muted p-3">
            <p className="text-xs font-medium text-muted-foreground">
              Net Savings
            </p>
            <p
              className={`text-lg font-bold ${
                totals.income - totals.expense >= 0
                  ? "text-emerald-600 dark:text-emerald-400"
                  : "text-rose-600 dark:text-rose-400"
              }`}
            >
              {formatCurrency(totals.income - totals.expense)}
            </p>
          </div>
        </div>

        {/* Chart */}
        <div className="h-[300px]">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart
              data={monthlyData}
              margin={{ top: 10, right: 10, left: 10, bottom: 0 }}
            >
              <CartesianGrid
                strokeDasharray="3 3"
                vertical={false}
                stroke="var(--color-border)"
              />
              <XAxis
                dataKey="monthLabel"
                axisLine={false}
                tickLine={false}
                tick={{ fontSize: 12, fill: "var(--color-muted-foreground)" }}
                tickFormatter={(value) => {
                  // Show abbreviated month on smaller screens
                  const parts = value.split(" ");
                  return parts[0].substring(0, 3);
                }}
              />
              <YAxis
                axisLine={false}
                tickLine={false}
                tick={{ fontSize: 12, fill: "var(--color-muted-foreground)" }}
                tickFormatter={(value) => {
                  if (value >= 1000) {
                    return `$${(value / 1000).toFixed(0)}k`;
                  }
                  return `$${value}`;
                }}
                width={50}
              />
              <Tooltip
                content={<MonthlyTooltip />}
                cursor={{ fill: "rgba(0, 0, 0, 0.05)" }}
              />
              <Legend
                formatter={(value) => (
                  <span className="text-sm capitalize text-foreground">
                    {value}
                  </span>
                )}
              />
              <Bar
                dataKey="income"
                name="Income"
                fill="#10b981"
                radius={[4, 4, 0, 0]}
                maxBarSize={40}
              />
              <Bar
                dataKey="expense"
                name="Expense"
                fill="#f43f5e"
                radius={[4, 4, 0, 0]}
                maxBarSize={40}
              />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}
