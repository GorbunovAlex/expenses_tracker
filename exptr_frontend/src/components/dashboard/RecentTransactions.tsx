"use client";

import * as React from "react";
import { format } from "date-fns";
import {
  ArrowDownLeft,
  ArrowUpRight,
  MoreHorizontal,
  Pencil,
  Trash2,
} from "lucide-react";

import { cn } from "@/lib/utils";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import type { Operation, Category } from "@/types/api";

interface RecentTransactionsProps {
  operations: Operation[];
  categories: Category[];
  limit?: number;
  onEdit?: (operation: Operation) => void;
  onDelete?: (operation: Operation) => void;
  className?: string;
}

export function RecentTransactions({
  operations,
  categories,
  limit = 10,
  onEdit,
  onDelete,
  className,
}: RecentTransactionsProps) {
  const getCategoryById = (id: string): Category | undefined => {
    return categories.find((cat) => cat.id === id);
  };

  const formatAmount = (amount: number, currency: string): string => {
    return new Intl.NumberFormat("en-US", {
      style: "currency",
      currency: currency || "USD",
      minimumFractionDigits: 0,
      maximumFractionDigits: 2,
    }).format(amount / 100); // Assuming amount is in cents
  };

  const sortedOperations = [...operations]
    .sort(
      (a, b) =>
        new Date(b.created_at).getTime() - new Date(a.created_at).getTime(),
    )
    .slice(0, limit);

  if (operations.length === 0) {
    return (
      <Card className={className}>
        <CardHeader>
          <CardTitle className="text-lg font-semibold">
            Recent Transactions
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex flex-col items-center justify-center py-12 text-center">
            <div className="rounded-full bg-muted p-4 mb-4">
              <ArrowUpRight className="h-8 w-8 text-muted-foreground" />
            </div>
            <h3 className="text-lg font-medium mb-2">No transactions yet</h3>
            <p className="text-sm text-muted-foreground max-w-sm">
              Start tracking your expenses by adding your first transaction.
            </p>
          </div>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card className={className}>
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle className="text-lg font-semibold">
          Recent Transactions
        </CardTitle>
        <Badge variant="outline" className="font-normal">
          {operations.length} total
        </Badge>
      </CardHeader>
      <CardContent className="p-0">
        <div className="divide-y divide-border">
          {sortedOperations.map((operation) => {
            const category = getCategoryById(operation.category_id);
            const isExpense = operation.type === "expense";

            return (
              <div
                key={operation.id}
                className="flex items-center justify-between px-6 py-4 hover:bg-muted/50 transition-colors"
              >
                {/* Left section: Icon + Details */}
                <div className="flex items-center gap-4">
                  {/* Transaction type icon */}
                  <div
                    className={cn(
                      "flex h-10 w-10 items-center justify-center rounded-full",
                      isExpense
                        ? "bg-destructive/10 text-destructive"
                        : "bg-emerald-500/10 text-emerald-500",
                    )}
                  >
                    {isExpense ? (
                      <ArrowDownLeft className="h-5 w-5" />
                    ) : (
                      <ArrowUpRight className="h-5 w-5" />
                    )}
                  </div>

                  {/* Transaction details */}
                  <div className="flex flex-col">
                    <span className="font-medium text-foreground">
                      {operation.name}
                    </span>
                    <div className="flex items-center gap-2 text-sm text-muted-foreground">
                      {category && (
                        <>
                          <span
                            className="inline-block h-2 w-2 rounded-full"
                            style={{ backgroundColor: category.color }}
                          />
                          <span>{category.name}</span>
                          <span>•</span>
                        </>
                      )}
                      <span>
                        {format(new Date(operation.created_at), "MMM d, yyyy")}
                      </span>
                    </div>
                  </div>
                </div>

                {/* Right section: Amount + Actions */}
                <div className="flex items-center gap-4">
                  <span
                    className={cn(
                      "font-semibold tabular-nums",
                      isExpense ? "text-destructive" : "text-emerald-500",
                    )}
                  >
                    {isExpense ? "-" : "+"}
                    {formatAmount(operation.amount, operation.currency)}
                  </span>

                  {(onEdit || onDelete) && (
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button
                          variant="ghost"
                          size="icon"
                          className="h-8 w-8 text-muted-foreground hover:text-foreground"
                        >
                          <MoreHorizontal className="h-4 w-4" />
                          <span className="sr-only">Open menu</span>
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        {onEdit && (
                          <DropdownMenuItem onClick={() => onEdit(operation)}>
                            <Pencil className="mr-2 h-4 w-4" />
                            Edit
                          </DropdownMenuItem>
                        )}
                        {onEdit && onDelete && <DropdownMenuSeparator />}
                        {onDelete && (
                          <DropdownMenuItem
                            onClick={() => onDelete(operation)}
                            className="text-destructive focus:text-destructive"
                          >
                            <Trash2 className="mr-2 h-4 w-4" />
                            Delete
                          </DropdownMenuItem>
                        )}
                      </DropdownMenuContent>
                    </DropdownMenu>
                  )}
                </div>
              </div>
            );
          })}
        </div>
      </CardContent>
    </Card>
  );
}
