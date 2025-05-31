package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goat",
	Short: "Go Application Toolkit",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, err := cmd.Flags().GetBool("version")
		if err != nil {
			fmt.Printf("Error parsing flags: %v", err)
			os.Exit(1)
		}

		if versionFlag {
			versionCmd.Run(cmd, args)
			os.Exit(0)
		}

		fmt.Println("Welcome to goat!\nTo get started, run \"goat new my-app\", or \"goat --help\"")

		err = cmd.Usage()
		if err != nil {
			fmt.Printf("Error printing usage: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print GOAT version")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
