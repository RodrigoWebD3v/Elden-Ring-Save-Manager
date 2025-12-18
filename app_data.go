package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func EldenRingPath() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(base, "EldenRing")

	return path, nil
}

func SavesPath() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(base, "saves")

	return path, nil
}

func CriarDiretorioSaves() bool {
	savesPath, err := SavesPath()
	if err != nil {
		panic(err)
	}

	if _, statErr := os.Stat(savesPath); os.IsNotExist(statErr) {
		if mkErr := os.MkdirAll(savesPath, os.ModePerm); mkErr != nil {
			panic(mkErr)
		}
		InicializarSaves(savesPath)
		return true
	} else if statErr != nil {
		panic(statErr)
	}

	InicializarSaves(savesPath)
	return false
}

func InicializarSaves(path string) bool {
	eldenRingPath, err := EldenRingPath()

	if err != nil {
		panic(err)
	}

	CopyDir(eldenRingPath, path)

	return true
}

func CriarConfigFile(path string, id string) bool {
	err := os.WriteFile(path+"/ativo.id", []byte(id), os.ModePerm)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de configuracoes: %s\n", err)
		return false
	}
	return true
}

func CriarIdFile(path string, id string) bool {
	err := os.WriteFile(path+"/save.id", []byte(id), os.ModePerm)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de ID: %s\n", err)
		return false
	}

	return true
}

func VerificaSaveAtivo(name string) bool {
	savesPath, err := SavesPath()

	if err != nil {
		panic(err)
	}
	saveIdPath := savesPath + "/" + name + "/save.id"
	if _, statErr := os.Stat(saveIdPath); os.IsNotExist(statErr) {
		return false
	}

	ativoAtualPath := savesPath + "/ativo.id"
	if _, statErr := os.Stat(ativoAtualPath); os.IsNotExist(statErr) {
		return false
	}

	idAtivoBruto, err := os.ReadFile(ativoAtualPath)
	if err != nil {
		panic(err)
	}

	idAtivoAtual := strings.TrimSpace(string(idAtivoBruto))

	idBruto, err := os.ReadFile(saveIdPath)
	if err != nil {
		panic(err)
	}

	id := strings.TrimSpace(string(idBruto))

	fmt.Printf("IDS SENDO COMPARADOS %s ==== %s", id, idAtivoAtual)
	switch id {
	case "":
		return false
	case idAtivoAtual:
		return true
	default:
		return false
	}
}
