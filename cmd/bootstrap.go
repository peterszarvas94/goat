package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func bootstrap(folderName string) error {
	path := folderName
	if path == "" {
		path = "."
	}

	if path != "." {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	output, err := cmd("git", "clone", "https://github.com/peterszarvas94/goat-bootstrap.git", path)
	fmt.Println(string(output))
	if err != nil {
		return err
	}

	name := path
	if name == "." {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		current := filepath.Base(wd)
		name = current
	}

	err = renameBootstrap(path, name)
	if err != nil {
		return err
	}

	err = os.Chdir(path)
	if err != nil {
		return err
	}

	err = makeEnv()
	if err != nil {
		return err
	}

	output, err = cmd("git", "remote", "remove", "origin")
	fmt.Println(string(output))
	if err != nil {
		return err
	}

	output, err = cmd("go", "install", "github.com/pressly/goose/v3/cmd/goose@latest")
	fmt.Println(output)
	if err != nil {
		return err
	}

	output, err = cmd("go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
	fmt.Println(output)
	if err != nil {
		return err
	}

	output, err = cmd("go", "install", "go install github.com/a-h/templ/cmd/templ@latest")
	fmt.Println(output)
	if err != nil {
		return err
	}

	err = migrateUpInitial()
	if err != nil {
		return err
	}

	fmt.Println("Default db schema is migrated")

	return nil
}

func renameBootstrap(path, name string) error {
	return filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		newContent := bytes.ReplaceAll(content, []byte("bootstrap"), []byte(name))

		return os.WriteFile(path, newContent, 0644)
	})
}

func makeEnv() error {
	_, err := os.Create(".env")
	if err != nil {
		return err
	}

	envContent := `DBPATH=sqlite.db
ENV=dev
	`
	return os.WriteFile(".env", []byte(envContent), 0655)
}
