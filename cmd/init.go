package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/peterszarvas94/goat/pkg/version"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:                   "init [name]",
	Short:                 "Initialize project",
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// parse target dir name and path

		var targetDir string
		if len(args) > 0 {
			targetDir = args[0]
		}

		if targetDir == "" || targetDir == "." {
			targetDir = "./"
		}

		fmt.Printf("Target dir is: %s\n", targetDir)

		targetDirFullPath, err := filepath.Abs(targetDir)
		if err != nil {
			fmt.Printf("Can not get target directory full path: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Target dir full path is: %s\n", targetDirFullPath)

		// parse tepmlate

		template, err := cmd.Flags().GetString("template")
		if err != nil {
			fmt.Printf("Can not parse flag \"--template\"%v", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Getting template: %s\n", template)

		// get project name

		projectName := targetDir
		if projectName == "./" {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err.Error())
			}
			projectName = filepath.Base(pwd)
		}

		fmt.Printf("Project name: %s\n", projectName)

		tmp, err := os.MkdirTemp("", "goat-template")
		if err != nil {
			fmt.Println("Error creating temp dir")
			os.Exit(1)
		}

		fmt.Printf("Temp dir created: %s\n", tmp)

		// clone repo

		err = utils.Cmd("git", "clone", "https://github.com/peterszarvas94/goat.git", tmp)
		if err != nil {
			fmt.Printf("Can not clone repo: %v", err)
			os.Exit(1)
		}

		err = os.Chdir(tmp)
		if err != nil {
			fmt.Printf("Can change directory to tmp: %v", err)
			os.Exit(1)
		}

		// checkout version

		_, err = utils.CmdWithOutput("git", "checkout", version.Version)
		if err != nil {
			fmt.Printf("Can checkout version: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Checked out version: %s\n", version.Version)

		err = utils.CopyDir(filepath.Join(tmp, "templates", template), targetDirFullPath)
		if err != nil {
			fmt.Printf("Can copy dir %s to %s: %v", tmp, targetDirFullPath, err)
			os.Exit(1)
		}

		// rename

		err = utils.ReplaceAllString(targetDirFullPath, template, projectName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Renamed %s to %s\n", template, projectName)

		// tidy

		err = utils.Cmd("go", "mod", "tidy")
		if err != nil {
			fmt.Printf("Error tidying: %v\n", err.Error())
			os.Exit(1)
		}

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

		err = os.Chdir(targetDirFullPath)
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

		err = utils.Cmd("go", "mod", "vendor")
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
