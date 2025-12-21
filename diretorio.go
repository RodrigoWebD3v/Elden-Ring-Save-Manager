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

func ConfigFile(id string, idSteam string) bool {
	savesPath, err := SavesPath()

	if _, err = os.Stat(savesPath + "/ativo.id"); err == nil {
		fmt.Println("ativo.id já existe")
		return true
	}

	content := fmt.Sprintf("%s;\n%s\n", id, idSteam)

	err = os.WriteFile(savesPath+"/ativo.id", []byte(content), 0644)
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
