package utils

import (
	"os"
	"path/filepath"
)

func GetSubfolders(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var subfolders []string
	for _, entry := range entries {
		if entry.IsDir() {
			subfolders = append(subfolders, filepath.Join(path, entry.Name()))
		}
	}

	return subfolders, nil
}
