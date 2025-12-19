import { useEffect, useState } from "react";
import { Header } from "@/components/Header";
import { SaveList } from "@/components/SaveList";
import { Save } from "@/components/SaveItem";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { FolderPlus, RefreshCw, HardDrive, Save as SaveIcon } from "lucide-react";
import { useToast } from "@/hooks/use-toast";
import { CarregarSaves, CriarBackup, Inicializar } from "../../wailsjs/go/main/App";

// Mock data para demonstracao
const mockSaves: Save[] = [
  {
    id: "1",
    name: "Tarnished_Principal",
    isActive: true,
    lastModified: "17 Dez 2024, 14:30",
    size: "45.2 MB",
  },
  {
    id: "2",
    name: "Backup_DLC_Shadow",
    isActive: false,
    lastModified: "15 Dez 2024, 22:15",
    size: "44.8 MB",
  },
  {
    id: "3",
    name: "Mage_Build_NG+",
    isActive: false,
    lastModified: "10 Dez 2024, 18:00",
    size: "43.1 MB",
  },
  {
    id: "4",
    name: "Strength_Build",
    isActive: false,
    lastModified: "05 Dez 2024, 09:45",
    size: "42.5 MB",
  },
];

const Index = () => {
  const [isInitialized, setIsInitialized] = useState(false);
  const [saves, setSaves] = useState<Save[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [newSaveName, setNewSaveName] = useState("");
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const { toast } = useToast();

  useEffect(() => {
    Inicializar().then((result) => {
      CarregarSaves().then((loadedSaves) => {
        setIsInitialized(true);
        setSaves(loadedSaves.map((save: any) => ({
          id: save.Id,
          name: save.Name,
          isActive: save.IsAtivo,
          lastModified: save.LastModified,
          size: save.Size,
        })));
      });

      toast({
        title: "Inicializado com sucesso!",
        description: "Pastas criadas e saves carregados.",
      });

    })
  }, []);

  const handleRefresh = async () => {
    setIsLoading(true);
    CarregarSaves().then((loadedSaves) => {
      setSaves(loadedSaves.map((save: any) => ({
        id: save.Id,
        name: save.Name,
        isActive: save.IsAtivo,
        lastModified: save.LastModified,
        size: save.Size,
      })));
    });
    setIsLoading(false);
    toast({
      title: "Lista atualizada",
      description: `${mockSaves.length} saves encontrados.`,
    });
  };

const handleCreateBackup = async () => {
  setIsLoading(true);
  CriarBackup(newSaveName.trim()).then((result) => {
    toast({
      title: "Backup criado!",
      description: `${newSaveName.trim()} salvo com sucesso.`,
    });

    if (result) {
      CarregarSaves().then((loadedSaves) => {
        setSaves(loadedSaves.map((save: any) => ({
          id: save.Id,
          name: save.Name,
          isActive: save.IsAtivo,
          lastModified: save.LastModified,
          size: save.Size,
        })));
        setIsLoading(false);
        setNewSaveName("");
        setIsDialogOpen(false);
      });
    }
  });
};

const handleActivate = (id: string) => {
  setSaves((prev) =>
    prev.map((save) => ({
      ...save,
      isActive: save.id === id,
    }))
  );
  const activatedSave = saves.find((s) => s.id === id);
  toast({
    title: "Save ativado",
    description: `${activatedSave?.name} agora e o save ativo.`,
  });
};

const handleDelete = (id: string) => {
  const saveToDelete = saves.find((s) => s.id === id);
  setSaves((prev) => prev.filter((save) => save.id !== id));
  toast({
    title: "Save excluido",
    description: `${saveToDelete?.name} foi removido.`,
    variant: "destructive",
  });
};

const handleDuplicate = (id: string) => {
  const saveToDuplicate = saves.find((s) => s.id === id);
  if (saveToDuplicate) {
    const newSave: Save = {
      ...saveToDuplicate,
      id: Date.now().toString(),
      name: `${saveToDuplicate.name}_copy`,
      isActive: false,
      lastModified: new Date().toLocaleString("pt-BR"),
    };
    setSaves((prev) => [...prev, newSave]);
    toast({
      title: "Save duplicado",
      description: `Copia de ${saveToDuplicate.name} criada.`,
    });
  }
};

return (
  <div className="min-h-screen bg-background p-6 md:p-8">
    <div className="max-w-2xl mx-auto">
      <Header />

      <Card className="border-border bg-card/50 backdrop-blur">
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
          <CardTitle className="text-lg font-display flex items-center gap-2">
            <HardDrive className="w-5 h-5 text-primary" />
            Saves
          </CardTitle>
          <div className="flex gap-2">
            <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
              <DialogTrigger asChild>
                <Button className="gap-2">
                  <SaveIcon className="w-4 h-4" />
                  Criar Backup
                </Button>
              </DialogTrigger>
              <DialogContent>
                <DialogHeader>
                  <DialogTitle className="font-display">Criar Backup do Save Atual</DialogTitle>
                  <DialogDescription>
                    Isso vai criar uma copia do save que esta rodando atualmente no jogo.
                  </DialogDescription>
                </DialogHeader>
                <div className="py-4">
                  <Input
                    placeholder="Nome do backup (ex: Boss_Malenia)"
                    value={newSaveName}
                    onChange={(e) => setNewSaveName(e.target.value)}
                    onKeyDown={(e) => e.key === "Enter" && handleCreateBackup()}
                  />
                </div>
                <DialogFooter>
                  <Button variant="outline" onClick={() => setIsDialogOpen(false)}>
                    Cancelar
                  </Button>
                  <Button onClick={handleCreateBackup} disabled={isLoading}>
                    {isLoading ? "Salvando..." : "Criar Backup"}
                  </Button>
                </DialogFooter>
              </DialogContent>
            </Dialog>
            <Button
              variant="outline"
              size="icon"
              onClick={handleRefresh}
              disabled={isLoading}
            >
              <RefreshCw className={`w-4 h-4 ${isLoading ? "animate-spin" : ""}`} />
            </Button>
          </div>
        </CardHeader>
        <CardContent>
          <SaveList
            saves={saves}
            onActivate={handleActivate}
            onDelete={handleDelete}
            onDuplicate={handleDuplicate}
          />
        </CardContent>
      </Card>

      <footer className="mt-8 text-center text-sm text-muted-foreground">
        <p>Elden Ring Save Manager v1.0</p>
      </footer>
    </div>
  </div>
);
};

export default Index;
