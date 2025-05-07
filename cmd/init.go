package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/dependencies"
	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:                   "init [name]",
	Short:                 "Initialize project",
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// parse dir name and template name

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

		// get project name

		projectName := targetDir
		if projectName == "./" {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err.Error())
			}
			projectName = filepath.Base(pwd)
		}

		fmt.Printf("Project name is %s\n", projectName)

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

		// init

		// err = utils.Cmd("go", "mod", "init", projectName)
		// if err != nil {
		// 	fmt.Println("Error initializing:", err.Error())
		// 	os.Exit(1)
		// }

		// rename

		err = utils.ReplaceAllString(targetDir, template, projectName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Renamed %s to %s\n", template, projectName)

		// install deps

		dependencies.InstallAll(template)

		// installs cli-s

		clis := []string{
			"github.com/pressly/goose/v3/cmd/goose@v3.24.2",
			"github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0",
			"github.com/a-h/templ/cmd/templ@v0.3.857",
		}

		for _, cli := range clis {
			err = utils.Cmd("go", "install", cli)
			if err != nil {
				fmt.Printf("Error installing %s: %v\n", cli, err.Error())
				os.Exit(1)
			}
		}

		// env

		err = os.Chdir(targetDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

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

		// git

		err = utils.Cmd("git", "init")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// tidy

		err = utils.Cmd("go", "mod", "tidy")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// vendor

		err = utils.Cmd("go", "mod", "vendor", template)
		if err != nil {
			fmt.Println("Error initializing:", err.Error())
			os.Exit(1)
		}

		// generate

		err = utils.Cmd("templ", "generate")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// db

		file, err = os.Create(constants.DBPath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()

		err = utils.CreateDirIfNotExists(constants.MigrationsDir)
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
			fmt.Println("No mogrations found")
			os.Exit(0)
		}

		err = utils.Cmd("goose", "-dir", constants.MigrationsDir, "sqlite3", constants.DBPath, "up")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Default db schema is migrated")
	},
}

func init() {
	initCmd.Flags().StringP("template", "t", "bare", "Specify a project template, e.g. \"bare\", \"basic-auth\"")
}
