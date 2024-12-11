package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/cmd/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goat",
	Short: "Go Application Toolkit",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to goat!\nTo get started, run \"goat new my-app\", or \"goat --help\"")
	},
}

var scaffholdCmd = &cobra.Command{
	Use:                   "new [name]",
	Short:                 "Scaffholfd project",
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		folderName := ""
		if len(args) > 0 {
			folderName = args[0]
		}

		err := commands.Scaffhold(folderName)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

var addModelCmd = &cobra.Command{
	Use:                   "model:add [name]",
	Short:                 "Add new model",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.ModelAdd(args[0])
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

var migrationNewCmd = &cobra.Command{
	Use:                   "mig:new [title]",
	Short:                 "Add new empty migration file",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.NewMigration(args[0])
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:                   "mig:up",
	Short:                 "Run up migrations",
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.MigrateUpDown("up")
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:                   "mig:down",
	Short:                 "Run one migration down",
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.MigrateUpDown("down")
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scaffholdCmd)
	rootCmd.AddCommand(addModelCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrationNewCmd)
}
