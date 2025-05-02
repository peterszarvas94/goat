package cmd

import (
	"fmt"

	"github.com/peterszarvas94/goat/pkg/utils"
	"github.com/spf13/cobra"
)

func generateModel() error {
	err := utils.Cmd("sqlc", "generate")
	return err
}

var modelGenCmd = &cobra.Command{
	Use:                   "model:gen [name]",
	Short:                 "Generate model from existing schemas",
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := generateModel()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}
