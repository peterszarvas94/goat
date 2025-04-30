package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/cmd/helpers"
	"github.com/peterszarvas94/goat/constants"
	"github.com/peterszarvas94/goat/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:                   "init [name]",
	Short:                 "Scaffhold project",
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// 0. parse dir name and template name

		var targetDir string
		if len(args) > 0 {
			targetDir = args[0]
		}

		if targetDir == "" || targetDir == "." {
			targetDir = "./"
		}

		template, err := cmd.Flags().GetString("template")
		if err != nil {
			fmt.Println(err.Error())
		}

		// 1. get project name

		projectName := targetDir
		if projectName == "./" {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err.Error())
			}
			projectName = filepath.Base(pwd)
		}

		fmt.Printf("Project name is %s\n", projectName)

		// 2. unzip

		err = utils.UnzipFromEmbed(embedZip, "tmp")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Unzipped to tmp\n")

		// 2. copy template

		templateDir := filepath.Join("tmp", template)
		err = utils.CopyDir(templateDir, targetDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Copied from tmp to %s\n", templateDir)

		err = os.RemoveAll("tmp")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("tmp removed")

		// 2. rename

		err = helpers.ReplaceAllString(targetDir, "scaffhold", projectName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Renamed %s to %s\n", "scaffhold", projectName)

		// 3. env

		err = os.Chdir(targetDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// TODO: include it in embedded repo
		file, err := os.Create(".env")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()

		fmt.Printf("Env file created\n")

		envContent := fmt.Sprintf(`DBPATH=%s
GOATENV=dev
PORT=9999
`,
			constants.DBPath,
		)

		_, err = file.Write([]byte(envContent))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Env file written\n")

		err = file.Chmod(0655)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// 4. git

		_, err = helpers.Cmd("git", "init")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// 5. installs

		_, err = helpers.Cmd("go", "install", "github.com/pressly/goose/v3/cmd/goose@latest")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, err = helpers.Cmd("go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, err = helpers.Cmd("go", "install", "github.com/a-h/templ/cmd/templ@latest")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, err = helpers.Cmd("go", "get", "github.com/a-h/templ/cmd/templ@latest")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, err = helpers.Cmd("go", "mod", "tidy")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, err = helpers.Cmd("templ", "generate")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// 6. db

		file, err = os.Create(constants.DBPath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()

		err = helpers.CreateDirIfNotExists(constants.MigrationsDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		entries, err := os.ReadDir(constants.MigrationsDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// no migrations found (for e.g. "bare" template)
		if len(entries) == 0 {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		_, err = helpers.Cmd("goose", "-dir", constants.MigrationsDir, "sqlite3", constants.DBPath, "up")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Default db schema is migrated")
	},
}
