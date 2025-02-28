package importmap

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/templ"
)

type ImportMap struct {
	Imports map[string]string `json:"imports"`
}

var Imports = ImportMap{}

func Setup(imports map[string]string) {
	Imports = ImportMap{
		Imports: imports,
	}
}

func ScriptTag() templ.Component {
	jsonData, _ := json.MarshalIndent(Imports, "", "  ")
	element := fmt.Sprintf(`<script type="importmap">
%s
</script>`, string(jsonData))

	return ImportMapComponent(element)
}

// func GenerateTSConfig() {
// }
