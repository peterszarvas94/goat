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
		fmt.Println("Welcome to the goat bootstrap!\nTo get started, run \"goat new my-app\", or \"goat --help\"")
	},
}

var bootstrapCmd = &cobra.Command{
	Use:                   "new [name]",
	Short:                 "Bootstrap project",
	Args:                  cobra.RangeArgs(0, 1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		folderName := ""
		if len(args) > 0 {
			folderName = args[0]
		}

		err := bootstrap(folderName)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

var addModelCmd = &cobra.Command{
	Use:                   "model add [name]",
	Short:                 "Add new model",
	Args:                  cobra.ExactArgs(2),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := modelAdd(args[1])
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
	rootCmd.AddCommand(bootstrapCmd)
	rootCmd.AddCommand(addModelCmd)
}
