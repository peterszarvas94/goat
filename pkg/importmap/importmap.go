package importmap

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/goat/pkg/constants"
	"github.com/peterszarvas94/goat/pkg/logger"
)

type ImportMap struct {
	Imports map[string]string `json:"imports"`
}

type CompilerOptions struct {
	CheckJs bool                `json:"checkJs"`
	AllowJs bool                `json:"allowJs"`
	NoEmit  bool                `json:"noEmit"`
	Lib     []string            `json:"lib"`
	BaseURL string              `json:"baseUrl"`
	Paths   map[string][]string `json:"paths"`
}

type TsConfig struct {
	CompilerOptions CompilerOptions `json:"compilerOptions"`
	Includes        []string        `json:"include"`
}

var scriptTag string = ""

// Loads "importmap.json" into "ScriptTag" templ component, which can be called in <head> element
//
// Generates "tsconfig.paths.json", which can be imported into "tsconfig.json" using "extends": "./tsconfig.paths.json"
func Setup() error {
	// read importmap.json
	file, err := os.ReadFile(constants.ImportMapFile)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	// store script tag as string
	scriptTag = fmt.Sprintf(`<script type="importmap">
%s
</script>`, string(file))

	// parse
	var importmap ImportMap
	if err := json.Unmarshal(file, &importmap); err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Debug("The importmap file is parsed")

	// make tsconfig struct
	tsConfigPaths := TsConfig{
		CompilerOptions: CompilerOptions{
			BaseURL: "./",
			CheckJs: true,
			AllowJs: true,
			NoEmit:  true,
			Lib:     []string{"esnext", "dom"},
			Paths:   make(map[string][]string),
		},
		Includes: []string{"./**/*.js"},
	}

	for key, val := range importmap.Imports {
		tsConfigPaths.CompilerOptions.Paths[key] = []string{strings.TrimPrefix(val, fmt.Sprintf("/%s/", constants.JSDir))}
	}

	// write tsconfig into file
	tsConfigPathsJSON, err := json.MarshalIndent(tsConfigPaths, "", "	")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	tsConfigPathsContent := fmt.Sprintf("%s\n%s", constants.DoNotModify, tsConfigPathsJSON)
	err = os.WriteFile(constants.TSConfigPahtsFile, []byte(tsConfigPathsContent), 0644)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Debug("The tsconfig file is written")

	return nil
}

// Renders importmap script tag
//
// Data should be loaded first from "importmap.json" by calling Setup
func ScriptTag() templ.Component {
	return ImportMapComponent(scriptTag)
}
