import { StatusIndicator } from "./StatusIndicator";
import { cn } from "@/lib/utils";
import { FileArchive, MoreVertical, Play, Trash2, Copy } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

export interface Save {
  id: string;
  name: string;
  isActive: boolean;
  lastModified: string;
  size: string;
}

interface SaveItemProps {
  save: Save;
  onActivate: (id: string) => void;
  onDelete: (id: string) => void;
  onDuplicate: (id: string) => void;
}

export const SaveItem = ({ save, onActivate, onDelete, onDuplicate }: SaveItemProps) => {
  return (
    <div
      className={cn(
        "group flex items-center gap-4 p-4 rounded-lg border transition-all duration-200 animate-fade-in",
        save.isActive
          ? "bg-secondary/80 border-primary/30 glow-gold"
          : "bg-card border-border hover:border-primary/20 hover:bg-secondary/50"
      )}
    >
      <StatusIndicator isActive={save.isActive} />
      
      <div className="flex items-center gap-3 flex-1 min-w-0">
        <FileArchive className="w-5 h-5 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className={cn(
            "font-medium truncate",
            save.isActive ? "text-primary" : "text-foreground"
          )}>
            {save.name}
          </p>
          <p className="text-sm text-muted-foreground">
            {save.lastModified} - {save.size}
          </p>
        </div>
      </div>

      <div className="flex items-center gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
        {!save.isActive && (
          <Button
            variant="ghost"
            size="sm"
            onClick={() => onActivate(save.id)}
            className="text-status-active hover:text-status-active hover:bg-status-active/10"
          >
            <Play className="w-4 h-4 mr-1" />
            Ativar
          </Button>
        )}
        
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" size="icon" className="h-8 w-8">
              <MoreVertical className="w-4 h-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem onClick={() => onDuplicate(save.id)}>
              <Copy className="w-4 h-4 mr-2" />
              Duplicar
            </DropdownMenuItem>
            <DropdownMenuItem 
              onClick={() => onDelete(save.id)}
              className="text-destructive focus:text-destructive"
            >
              <Trash2 className="w-4 h-4 mr-2" />
              Excluir
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </div>
  );
};
