package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/spf13/cobra"
)

// exampleCmd represents the example command
var exampleCmd = &cobra.Command{
	Use:     "examples",
	Aliases: []string{"e"},
	Short:   "List all available example templates",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		subfolders, err := utils.GetSubfolders(constants.ExamplesDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Avaiable examlpes:")
		for folder := range slices.Values(subfolders) {
			fmt.Printf("- %s\n", strings.Split(folder, "/")[1])
		}

	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}
