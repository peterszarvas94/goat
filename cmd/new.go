package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/cmd/helpers"
	"github.com/peterszarvas94/goat/constants"
	"github.com/peterszarvas94/goat/files"
	"github.com/spf13/cobra"
)

func initializeProject(folderName, template string) error {
	/* 1. PATH */

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

	/* 2. DOWNLOAD */
	// TODO: embed this, don't download

	// make tmp dir
	err := os.MkdirAll("tmp", 0755)
	if err != nil {
		return err
	}

	// downloading the repo
	_, err = helpers.Cmd("git", "clone", "https://github.com/peterszarvas94/goat-scaffhold.git", "tmp")
	if err != nil {
		return err
	}

	// if current repo, get pwd name
	name := path
	if name == "." {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		current := filepath.Base(wd)
		name = current
	}

	// copy the template files
	templateDir := filepath.Join("tmp", "templates", template)

	err = files.CopyDir(templateDir, path)
	if err != nil {
		return err
	}

	// delete tmp dir
	err = os.RemoveAll("tmp")
	if err != nil {
		return err
	}

	// rename project
	err = helpers.ReplaceAll(path, "scaffhold", name)
	if err != nil {
		return err
	}

	err = os.Chdir(path)
	if err != nil {
		return err
	}

	/* 3. ENV */
	// TODO: include it in embedded repo

	file, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	envContent := fmt.Sprintf(`DBPATH=%s
GOATENV=dev
PORT=9999
`,
		constants.DBPath,
	)

	_, err = file.Write([]byte(envContent))
	if err != nil {
		return err
	}

	err = file.Chmod(0655)
	if err != nil {
		return err
	}

	/* 4. GIT */

	_, err = helpers.Cmd("git", "init")
	if err != nil {
		return err
	}

	/* 5. INSTALLS */

	_, err = helpers.Cmd("go", "install", "github.com/pressly/goose/v3/cmd/goose@latest")
	if err != nil {
		return err
	}

	_, err = helpers.Cmd("go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
	if err != nil {
		return err
	}

	/* 6. TEMPL */

	_, err = helpers.Cmd("go", "install", "github.com/a-h/templ/cmd/templ@latest")
	if err != nil {
		return err
	}

	_, err = helpers.Cmd("go", "get", "github.com/a-h/templ/cmd/templ@latest")
	if err != nil {
		return err
	}

	_, err = helpers.Cmd("go", "mod", "tidy")
	if err != nil {
		return err
	}

	_, err = helpers.Cmd("templ", "generate")
	if err != nil {
		return err
	}

	/* 7. DB */

	file, err = os.Create(constants.DBPath)
	if err != nil {
		return err
	}
	defer file.Close()

	/* 8. MIGRATIONS */

	err = helpers.CreateDirIfNotExists(constants.MigrationsDir)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(constants.MigrationsDir)
	if err != nil {
		return err
	}

	// no migrations found (for e.g. "bare" template)
	if len(entries) == 0 {
		return nil
	}

	_, err = helpers.Cmd("goose", "-dir", constants.MigrationsDir, "sqlite3", constants.DBPath, "up")
	if err != nil {
		return err
	}

	fmt.Println("Default db schema is migrated")

	return nil
}

var initCmd = &cobra.Command{
	Use:                   "init [name]",
	Short:                 "Scaffhold project",
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		folderName := ""
		if len(args) > 0 {
			folderName = args[0]
		}

		template, err := cmd.Flags().GetString("template")
		if template == "" || err != nil {
			fmt.Println(err.Error())
			template = "bare"
		}

		err = initializeProject(folderName, template)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}
