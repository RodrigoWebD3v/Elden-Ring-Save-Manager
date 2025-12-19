package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyDirInitialization(src string, dst string) error {
	return copyDir(src, dst, true, "")
}

func CopyDir(src string, dst string, name string) error {
	return copyDir(src, dst, true, name)
}

// copyDir copia recursivamente e só grava o arquivo de config na chamada raiz.
func copyDir(src string, dst string, isRoot bool, name string) error {
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

		nome := name

		if name == "" {
			nome = entry.Name()
		}

		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, nome)

		if entry.IsDir() {
			// Recursão
			if err := copyDir(srcPath, dstPath, false, ""); err != nil {
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
		CriarConfigFile(dst, id, name)
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
