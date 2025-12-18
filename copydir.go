package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyDir(src string, dst string) error {
	return copyDir(src, dst, true)
}

// copyDir copia recursivamente e só grava o arquivo de config na chamada raiz.
func copyDir(src string, dst string, isRoot bool) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	id := GerarId()

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursão
			if err := copyDir(srcPath, dstPath, false); err != nil {
				return err
			}

			CriarIdFile(dstPath, id)

		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	if isRoot {
		fmt.Printf("%s\n", dst)
		CriarConfigFile(dst, id)
	}

	return nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return out.Sync()
}

func TornarAtivo(name string) bool {
	savesPath, err := SavesPath()
	if err != nil {
		panic(err)
	}
	savePath := filepath.Join(savesPath, name)

	//msg := fmt.Sprintf("Ativo %s", "true")

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
