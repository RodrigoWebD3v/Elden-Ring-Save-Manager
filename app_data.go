package main

import (
	"fmt"
	"os"
	"strings"
)

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
