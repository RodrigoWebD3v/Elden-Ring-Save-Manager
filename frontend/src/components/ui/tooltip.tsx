import React from "react";

// Simplified no-op tooltip components to avoid build issues with Radix + JSX.
// They keep the API surface so other components don't break, but render plainly.
type TooltipChildren = { children: React.ReactNode };

const TooltipProvider: React.FC<TooltipChildren> = ({ children }) => <>{children}</>;

const Tooltip: React.FC<TooltipChildren> = ({ children }) => <>{children}</>;

const TooltipTrigger = React.forwardRef<HTMLSpanElement, React.HTMLAttributes<HTMLSpanElement>>(
  ({ children, ...props }, ref) => (
    <span ref={ref} {...props}>
      {children}
    </span>
  ),
);
TooltipTrigger.displayName = "TooltipTrigger";

const TooltipContent = React.forwardRef<HTMLDivElement, React.HTMLAttributes<HTMLDivElement>>(
  ({ children, ...props }, ref) => (
    <div ref={ref} style={{ display: "none" }} {...props}>
      {children}
    </div>
  ),
);
TooltipContent.displayName = "TooltipContent";

export { Tooltip, TooltipTrigger, TooltipContent, TooltipProvider };
