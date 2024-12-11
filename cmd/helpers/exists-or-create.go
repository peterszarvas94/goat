package helpers

import (
	"errors"
	"fmt"
	"os"
)

func ExistsOrCreateDir(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// if not existing, create
		return os.MkdirAll(path, 0755)
	}

	// if exists, no error
	return nil
}

func ExistsOrCreateFile(path string) (string, error) {
	// Check if the file exists
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// If the file does not exist, create it
		_, err = os.Create(path)
		if err != nil {
			return "", err
		}
		return "", nil // Return an empty string for a newly created file
	}

	// If the file exists, read its content
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func createFileIfNotExists(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// if no existing, create
		_, err = os.Create(path)
		return err
	}

	// if existst, error
	return fmt.Errorf("file %s already exists", path)
}
