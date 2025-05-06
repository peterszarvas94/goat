package cmd

import (
	"fmt"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GOAT version",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constants.Version)
	},
}
