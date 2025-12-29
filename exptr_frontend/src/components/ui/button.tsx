import * as React from "react";
import { Slot } from "@radix-ui/react-slot";
import { cva, type VariantProps } from "class-variance-authority";

import { cn } from "@/lib/utils";

const buttonVariants = cva(
  "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive cursor-pointer",
  {
    variants: {
      variant: {
        // Primary: Glacial Teal background, white text, hover to Deep Cerulean
        default:
          "bg-primary text-primary-foreground shadow-sm hover:bg-secondary",
        // Destructive
        destructive:
          "bg-destructive text-destructive-foreground shadow-sm hover:bg-destructive/90",
        // Outline: Transparent with Glacial Teal border
        outline:
          "border border-primary bg-transparent text-primary shadow-sm hover:bg-primary hover:text-primary-foreground",
        // Secondary: Deep Cerulean background
        secondary:
          "bg-secondary text-secondary-foreground shadow-sm hover:bg-primary",
        // Ghost: No background, teal text
        ghost: "text-primary hover:bg-primary/10 hover:text-secondary",
        // Link style
        link: "text-secondary underline-offset-4 hover:underline hover:text-primary",
        // Gradient: Hero gradient from Glacial Teal to Deep Cerulean
        gradient:
          "bg-gradient-to-br from-glacial-teal to-deep-cerulean text-white shadow-md hover:shadow-lg hover:from-deep-cerulean hover:to-glacial-teal",
      },
      size: {
        default: "h-9 px-4 py-2 has-[>svg]:px-3",
        sm: "h-8 rounded-md gap-1.5 px-3 has-[>svg]:px-2.5 text-xs",
        lg: "h-10 rounded-md px-6 has-[>svg]:px-4",
        xl: "h-12 rounded-md px-8 has-[>svg]:px-5 text-base",
        icon: "size-9",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  }
);

function Button({
  className,
  variant,
  size,
  asChild = false,
  ...props
}: React.ComponentProps<"button"> &
  VariantProps<typeof buttonVariants> & {
    asChild?: boolean;
  }) {
  const Comp = asChild ? Slot : "button";

  return (
    <Comp
      data-slot="button"
      className={cn(buttonVariants({ variant, size, className }))}
      {...props}
    />
  );
}

export { Button, buttonVariants };
