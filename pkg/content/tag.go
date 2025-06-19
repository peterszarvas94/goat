package content

import (
	"fmt"
	"maps"
	"slices"
)

func GetAllTaggedFiles() FileMap {
	var res = make(FileMap)
	for route, file := range maps.All(Files) {
		for key := range maps.Keys(file.Frontmatter) {
			if key == "tag" {
				res[route] = file
			}
		}
	}
	return res
}

func GetTaggedFiles(tag string) FileMap {
	var res = make(FileMap)

	for route, file := range maps.All(Files) {
	inner:
		for key, value := range maps.All(file.Frontmatter) {
			if key != "tag" {
				continue inner
			}

			tags, ok := value.([]any)
			if !ok {
				fmt.Printf("The 'tag' value in frontmatter is not a slice in file '%s'\n", file.MarkdownPath)
				continue inner
			}

			for t := range slices.Values(tags) {
				err, parsed := parseTag(t)
				if err != nil {
					fmt.Printf("Error parsing tag '%v' for file '%s':\n%v", t, file.MarkdownPath, err)
					continue inner
				}
				if tag == parsed {
					res[route] = file
				}
			}
		}
	}
	return res
}

func parseTag(tag any) (error, string) {
	parsed, ok := tag.(string)
	if !ok {
		return fmt.Errorf("Cannot parse tag %v, type of %T\n", tag, tag), ""
	}
	return nil, parsed
}
