package main

import (
	"context"
	"fmt"
	"os"
)

// Save struct
type Save struct {
	Id           int
	Name         string
	IsAtivo      bool
	LastModified string
	Size         string
}

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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time! a", name)
}

func (a *App) ListaDiretorios() string {
	diretorio, err := EldenRingPath()

	if err != nil {
		panic(err)
	}

	entries, err := os.ReadDir(diretorio)

	if err != nil {
		panic(err)
	}

	return entries[0].Name()
}

func (a *App) Inicializar() bool {
	CriarDiretorioSaves()
	return true
}

func (a *App) CarregarSaves() []Save {

	savesPath, err := SavesPath()

	var saves []Save

	if err != nil {
		fmt.Printf("Erro ao executar CarregarSaves")
		panic(err)
	}

	GerarId()

	entries, err := os.ReadDir(savesPath)

	if err != nil {
		fmt.Printf("Erro ao ler diretorio de backup\n")
	}

	for _, entry := range entries {
		info, err := entry.Info()

		if err != nil {
			fmt.Printf("Erro ao ler informacoes do save")
			panic(err)
		}

		if entry.Name() == "ativo.id" {
			continue
		}

		save := Save{
			Name:         entry.Name(),
			IsAtivo:      VerificaSaveAtivo(entry.Name()),
			LastModified: info.ModTime().String(),
			Size:         fmt.Sprintf("%d bytes", info.Size()),
		}

		saves = append(saves, save)
	}

	return saves
}

func (a *App) DeletarDiretorios() bool {
	savesPath, err := SavesPath()

	if err != nil {
		fmt.Printf("Erro ao executar DeletarDiretorios")
		panic(err)
	}

	err = os.RemoveAll(savesPath)

	if err != nil {
		fmt.Printf("Erro ao deletar diretorio de backup")
		panic(err)
	}

	return true
}
