package cmd

import (
	"os"
)

func mkdir(name string) error {
	dirName := name
	err := os.Mkdir(dirName, 0755)
	return err
}
