package utils

import (
	"fmt"
	"os"
)

func CreateDirIfNotExists(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func CreateNonExistingFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsExist(err) {
			return nil, fmt.Errorf("File %s already exists", path)
		}
		return nil, err
	}
	return file, nil
}
