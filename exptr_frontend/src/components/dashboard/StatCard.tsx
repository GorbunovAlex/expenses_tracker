import { cn } from "@/lib/utils";
import { Card, CardContent } from "@/components/ui/card";
import { LucideIcon } from "lucide-react";

interface StatCardProps {
  title: string;
  value: string | number;
  subtitle?: string;
  icon?: LucideIcon;
  trend?: {
    value: number;
    isPositive: boolean;
  };
  variant?: "default" | "income" | "expense" | "balance";
  className?: string;
}

export function StatCard({
  title,
  value,
  subtitle,
  icon: Icon,
  trend,
  variant = "default",
  className,
}: StatCardProps) {
  const variantStyles = {
    default: {
      iconBg: "bg-muted",
      iconColor: "text-secondary",
    },
    income: {
      iconBg: "bg-emerald-100 dark:bg-emerald-900/30",
      iconColor: "text-emerald-600 dark:text-emerald-400",
    },
    expense: {
      iconBg: "bg-rose-100 dark:bg-rose-900/30",
      iconColor: "text-rose-600 dark:text-rose-400",
    },
    balance: {
      iconBg: "bg-primary/10",
      iconColor: "text-primary",
    },
  };

  const styles = variantStyles[variant];

  return (
    <Card className={cn("relative overflow-hidden", className)}>
      <CardContent className="p-6">
        <div className="flex items-start justify-between">
          <div className="space-y-2">
            <p className="text-sm font-medium text-muted-foreground">{title}</p>
            <div className="flex items-baseline gap-2">
              <h3 className="text-2xl font-bold tracking-tight">{value}</h3>
              {trend && (
                <span
                  className={cn(
                    "text-xs font-medium",
                    trend.isPositive
                      ? "text-emerald-600 dark:text-emerald-400"
                      : "text-rose-600 dark:text-rose-400"
                  )}
                >
                  {trend.isPositive ? "+" : ""}
                  {trend.value}%
                </span>
              )}
            </div>
            {subtitle && (
              <p className="text-xs text-muted-foreground">{subtitle}</p>
            )}
          </div>
          {Icon && (
            <div
              className={cn(
                "flex h-12 w-12 items-center justify-center rounded-lg",
                styles.iconBg
              )}
            >
              <Icon className={cn("h-6 w-6", styles.iconColor)} />
            </div>
          )}
        </div>
      </CardContent>
      {/* Decorative gradient line at bottom */}
      {variant !== "default" && (
        <div
          className={cn(
            "absolute bottom-0 left-0 right-0 h-1",
            variant === "income" && "bg-gradient-to-r from-emerald-400 to-emerald-600",
            variant === "expense" && "bg-gradient-to-r from-rose-400 to-rose-600",
            variant === "balance" && "bg-gradient-to-r from-glacial-teal to-deep-cerulean"
          )}
        />
      )}
    </Card>
  );
}
