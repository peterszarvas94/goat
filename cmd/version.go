package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GOAT version",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		version, err := utils.GetVersion()
		if err != nil {
			fmt.Printf("Can not get version: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println(version)
	},
}
