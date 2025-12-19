package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Save struct {
	Id           int
	Name         string
	IsAtivo      bool
	LastModified string
	Size         string
}

func CarregarSaves() []Save {

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

		fmt.Printf("Busca Saves\n")

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

func DeletarDiretorios() bool {
	savesPath, err := SavesPath()

	if err != nil {
		fmt.Printf("Erro ao executar DeletarDiretorios")
		return false
	}

	err = os.RemoveAll(savesPath)

	if err != nil {
		fmt.Printf("Erro ao deletar diretorio de backup")
		return false
	}

	return true
}

func TornarAtivo(name string) bool {
	savesPath, err := SavesPath()
	if err != nil {
		panic(err)
	}
	savePath := filepath.Join(savesPath, name)

	file, err := os.Open(savePath + "/save.id")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)

	linhaAtivo := ""
	for s.Scan() {
		linhaAtivo = s.Text()
		fmt.Printf("Tornar Ativo %s", linhaAtivo)
	}

	return true
}