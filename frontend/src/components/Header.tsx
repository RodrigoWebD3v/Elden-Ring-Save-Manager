import { Sword } from "lucide-react";

export const Header = () => {
  return (
    <header className="flex items-center gap-3 mb-8">
      <div className="p-3 rounded-lg bg-primary/10 border border-primary/20">
        <Sword className="w-8 h-8 text-primary" />
      </div>
      <div>
        <h1 className="text-2xl font-display font-semibold text-gradient-gold">
          Elden Ring Save Manager
        </h1>
        <p className="text-sm text-muted-foreground">
          Gerencie seus saves de forma simples e segura
        </p>
      </div>
    </header>
  );
};
