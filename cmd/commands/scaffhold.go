package commands

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/cmd/helpers"
)

func Scaffhold(folderName string) error {
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

	output, err := helpers.Cmd("git", "clone", "https://github.com/peterszarvas94/goat-scaffhold.git", path)
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

	err = renameScaffhold(path, name)
	if err != nil {
		return err
	}

	err = os.Chdir(path)
	if err != nil {
		return err
	}

	_, err = os.Create("sqlite.db")
	if err != nil {
		return err
	}

	_, err = GenerateMigration("create", "user")
	if err != nil {
		return err
	}

	_, err = GenerateMigration("create", "session")
	if err != nil {
		return err
	}

	err = makeEnv()
	if err != nil {
		return err
	}

	output, err = helpers.Cmd("git", "remote", "remove", "origin")
	fmt.Println(string(output))
	if err != nil {
		return err
	}

	output, err = helpers.Cmd("go", "install", "github.com/pressly/goose/v3/helpers.Cmd/goose@latest")
	fmt.Println(output)
	if err != nil {
		return err
	}

	output, err = helpers.Cmd("go", "install", "github.com/sqlc-dev/sqlc/helpers.Cmd/sqlc@latest")
	fmt.Println(output)
	if err != nil {
		return err
	}

	output, err = helpers.Cmd("go", "install", "github.com/a-h/templ/helpers.Cmd/templ@latest")
	fmt.Println(output)
	if err != nil {
		return err
	}

	output, err = helpers.Cmd("go", "install", "ariga.io/atlas-go-sdk/atlasexec@latest")

	fmt.Println(output)
	if err != nil {
		return err
	}

	output, err = helpers.Cmd("templ", "generate")
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

func renameScaffhold(path, name string) error {
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

		newContent := bytes.ReplaceAll(content, []byte("scaffhold"), []byte(name))

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