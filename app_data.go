package main

import (
	"os"
	"strings"
)

func VerificaSaveAtivo(name string) bool {
	savesPath, err := SavesPath()

	if err != nil {
		panic(err)
	}
	saveBkpIdPath := savesPath + "/" + name + "/save.id"
	if _, statErr := os.Stat(saveBkpIdPath); os.IsNotExist(statErr) {
		return false
	}

	ativoSaves := savesPath + "/ativo.id"
	if _, statErr := os.Stat(ativoSaves); os.IsNotExist(statErr) {
		return false
	}

	idAtivoSavesBytes, err := os.ReadFile(ativoSaves)
	if err != nil {
		panic(err)
	}

	idAtivoSaves := []byte(strings.Split(string(idAtivoSavesBytes), ";")[0])

	idAtivoAtual := strings.TrimSpace(string(idAtivoSaves))

	idSaveBkp, err := os.ReadFile(saveBkpIdPath)

	if err != nil {
		panic(err)
	}

	id := strings.TrimSpace(string(idSaveBkp))

	switch id {
	case "":
		return false
	case idAtivoAtual:
		return true
	default:
		return false
	}
}
