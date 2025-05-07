package dependencies

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/goat/pkg/utils"
)

func InstallAll(template string) {

	version, err := utils.GetVersion()
	if err != nil {
		fmt.Printf("Can not get version: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println(version)

	templates := map[string][]string{
		"bare": {
			"github.com/a-h/templ/cmd/templ@v0.3.857",
			fmt.Sprintf("github.com/peterszarvas94/goat@%s", version),
		},
		"basic-auth": {
			"github.com/a-h/templ/cmd/templ@v0.3.857",
			fmt.Sprintf("github.com/peterszarvas94/goat@%s", version),
		},
	}

	projectTemplate := templates[template]

	for _, dependency := range projectTemplate {
		err := utils.Cmd("go", "get", "-u", dependency)
		if err != nil {
			fmt.Printf("Error installing %s: %v\n", dependency, err.Error())
			os.Exit(1)
		}
	}

	fmt.Sprintln("Dependencies installed")
}
