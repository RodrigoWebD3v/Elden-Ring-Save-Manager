import { Save, SaveItem } from "./SaveItem";
import { ScrollArea } from "@/components/ui/scroll-area";
import { FolderOpen } from "lucide-react";

interface SaveListProps {
  saves: Save[];
  onActivate: (id: string) => void;
  onDelete: (id: string) => void;
  onDuplicate: (id: string) => void;
}

export const SaveList = ({ saves, onActivate, onDelete, onDuplicate }: SaveListProps) => {
  if (saves.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
        <FolderOpen className="w-16 h-16 mb-4 opacity-50" />
        <p className="text-lg font-medium">Nenhum save encontrado</p>
        <p className="text-sm">Inicialize o gerenciador para comecar</p>
      </div>
    );
  }

  return (
    <ScrollArea className="h-[400px] pr-4">
      <div className="space-y-3">
        {saves.map((save) => (
          <SaveItem
            key={save.id}
            save={save}
            onActivate={onActivate}
            onDelete={onDelete}
            onDuplicate={onDuplicate}
          />
        ))}
      </div>
    </ScrollArea>
  );
};
