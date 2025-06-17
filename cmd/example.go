package cmd

import (
	"fmt"
	"slices"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/spf13/cobra"
)

// exampleCmd represents the example command
var exampleCmd = &cobra.Command{
	Use:     "examples",
	Aliases: []string{"e"},
	Short:   "List all available example templates",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Avaiable examlpes:")
		for folder := range slices.Values(constants.Examples) {
			fmt.Printf("- %s\n", folder)
		}

	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}
