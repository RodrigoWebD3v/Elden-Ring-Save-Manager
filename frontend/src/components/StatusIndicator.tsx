import { cn } from "@/lib/utils";

interface StatusIndicatorProps {
  isActive: boolean;
  size?: "sm" | "md" | "lg";
}

export const StatusIndicator = ({ isActive, size = "md" }: StatusIndicatorProps) => {
  const sizeClasses = {
    sm: "w-2 h-2",
    md: "w-3 h-3",
    lg: "w-4 h-4",
  };

  return (
    <div
      className={cn(
        "rounded-full transition-all duration-300",
        sizeClasses[size],
        isActive
          ? "bg-status-active glow-green animate-pulse-glow"
          : "bg-status-inactive"
      )}
    />
  );
};
