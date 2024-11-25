package cmd

import (
	"errors"
	"fmt"
	"os"
)

func existOrCreateDir(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// if not existing, create
		return os.MkdirAll(path, 0755)
	}

	// if exists, no error
	return nil
}

func existsOrCreateFile(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// if no existing, create
		_, err = os.Create(path)
		return err
	}

	// if existst, error
	return fmt.Errorf("file %s already exists", path)
}
