package dependencies

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/utils"
)

var Templates = map[string][]string{
	"bare": {
		"github.com/a-h/templ/cmd/templ@v0.3.857",
		fmt.Sprintf("github.com/peterszarvas94/goat@%s", constants.Version),
	},
	"basic-auth": {
		"github.com/a-h/templ/cmd/templ@v0.3.857",
		fmt.Sprintf("github.com/peterszarvas94/goat@%s", constants.Version),
	},
}

func InstallAll(template string) {
	projectTemplate := Templates[template]

	for _, dependency := range projectTemplate {
		err := utils.Cmd("go", "get", "-u", dependency)
		if err != nil {
			fmt.Printf("Error installing %s: %v\n", dependency, err.Error())
			os.Exit(1)
		}
	}

	fmt.Sprintln("Dependencies installed")
}
