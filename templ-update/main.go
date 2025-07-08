package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/utils"
)

func main() {
	// install templ
	cli := fmt.Sprintf("github.com/a-h/templ/cmd/templ@%s", constants.TemplVersion)
	err := utils.Cmd("go", "install", cli)
	if err != nil {
		fmt.Printf("Error installing templ: %v\n", err)
		os.Exit(1)
	}

	// update templ for examples
	subfolders, err := utils.GetSubfolders(constants.ExamplesDir)
	if err != nil {
		fmt.Printf("Can not get subfolders: %v\n", err)
		os.Exit(1)
	}

	for _, subfolder := range subfolders {
		fmt.Printf("Updating templ in %s\n", subfolder)

		// Change to subfolder directory
		err = os.Chdir(filepath.Join(".", subfolder))
		if err != nil {
			fmt.Printf("Error changing to directory %s: %v\n", subfolder, err)
			os.Exit(1)
		}

		cli := fmt.Sprintf("github.com/a-h/templ@%s", constants.TemplVersion)
		err = utils.Cmd("go", "get", "-u", cli)
		if err != nil {
			fmt.Printf("Error installing %s in %s: %v\n", cli, subfolder, err.Error())
			os.Exit(1)
		}

		// Change back to parent directory
		err = os.Chdir(filepath.Join("..", ".."))
		if err != nil {
			fmt.Printf("Error changing back to parent directory: %v\n", err)
			os.Exit(1)
		}
	}
}
