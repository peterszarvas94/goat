package cmd

import (
	"fmt"

	"github.com/peterszarvas94/goat/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print GOAT version",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Get())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
