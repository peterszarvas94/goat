package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed embed.zip
var embedZip []byte

var rootCmd = &cobra.Command{
	Use:   "goat",
	Short: "Go Application Toolkit",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to goat!\nTo get started, run \"goat new my-app\", or \"goat --help\"")
	},
}

func init() {
	initCmd.Flags().StringP("template", "t", "bare", "Specify a project template, e.g. \"bare\", \"basic-auth\"")
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(modelAddCmd)
	rootCmd.AddCommand(modelGenCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrationNewCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
