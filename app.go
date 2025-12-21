package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Inicializar() bool {
	return CriarDiretorioSaves()
}

func (a *App) CarregarSaves() []Save {
	return CarregarSaves()
}

func (a *App) CriarBackup(name string) bool {
	return CriarBackupSave(name)
}

func (a *App) TornarAtivo(name string, id string) bool {
	return TornarAtivo(name, id)
}

func (a *App) DeletarDiretorios() bool {
	return DeletarDiretorios()
}
