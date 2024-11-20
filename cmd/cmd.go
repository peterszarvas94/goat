package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goat",
	Short: "GOAT bootstrap",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the goat bootstrap!")
		fmt.Println("To get started, run \"goat new my-app\"")
	},
}

var newCmd = &cobra.Command{
	Use:                   "new [name]",
	Short:                 "Create a new GOAT project",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := mkdir(args[0])
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		fmt.Printf("Directory '%s' created successfully!\n", args[0])
	},
}

var linkCmd = &cobra.Command{
	Use:                   "link [path]",
	Short:                 "Link local project",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := linkLocal(args[0])
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
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(linkCmd)
}
