package main

import (
	"os"
	"path/filepath"
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
