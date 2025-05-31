package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:                   "mig:down",
	Aliases:               []string{"md"},
	Short:                 "Run one migration down",
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := migrate("down")
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateDownCmd)
}
