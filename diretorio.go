package main

import (
	"fmt"
	"os"
)

func CriarDiretorioSaves() bool {
	savesPath, err := SavesPath()
	if err != nil {
		panic(err)
	}

	if _, statErr := os.Stat(savesPath); os.IsExist(statErr) {
		return true
	} else {
		if mkErr := os.MkdirAll(savesPath, os.ModePerm); mkErr != nil {
			panic(mkErr)
		}
		return true
	}
}

func CriarBackupSave(name string) bool {
	eldenRingPath, err := EldenRingPath()

	if err != nil {
		return false
	}

	savesPath, err := SavesPath()

	if err != nil {
		return false
	}

	CopyDir(eldenRingPath, savesPath, name)

	return true
}

func CriarConfigFile(path string, id string, idSteam string) bool {
	if _, statErr := os.Stat(path + "/ativo.id"); os.IsExist(statErr) {
		fmt.Printf("ativo.id ja existe")
		return true
	}

	content := fmt.Sprintf("%s;\n%s\n", id, idSteam)

	err := os.WriteFile(path+"/ativo.id", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de configuracoes: %s\n", err)
		return false
	}
	return true
}

func CriarIdFile(path string, id string) bool {
	err := os.WriteFile(path+"/save.id", []byte(id), 0644)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de ID: %s\n", err)
		return false
	}

	return true
}
