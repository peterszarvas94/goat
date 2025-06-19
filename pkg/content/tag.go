package content

import (
	"fmt"
	"maps"
	"slices"
)

func GetTaggedFiles(tag string) FileMap {
	var res = make(FileMap)

	for route, file := range maps.All(Files) {
		tags, err := GetTags(file)
		if err != nil {
			fmt.Printf("Could not get tags from file '%s':\n%v", file.MarkdownPath, err)
			return make(FileMap)
		}
		if slices.Contains(tags, tag) {
			res[route] = file
		}
	}
	return res
}

func parseTag(tag any) (string, error) {
	parsed, ok := tag.(string)
	if !ok {
		return "", fmt.Errorf("Cannot parse tag %v, type of %T\n", tag, tag)
	}
	return parsed, nil
}

func GetTags(file File) ([]string, error) {
	res := []string{}

	for key, value := range maps.All(file.Frontmatter) {
		if key != "tag" {
			continue
		}

		tags, ok := value.([]any)
		if !ok {
			return []string{}, fmt.Errorf("The 'tag' value in frontmatter is not a slice in file '%s'\n", file.MarkdownPath)
		}

		for t := range slices.Values(tags) {
			parsed, err := parseTag(t)
			if err != nil {
				return []string{}, fmt.Errorf("Error parsing tag '%v' for file '%s':\n%v", t, file.MarkdownPath, err)
			}

			res = append(res, parsed)
		}
	}

	return res, nil
}

func GetTagsSafe(file File) []string {
	tags, _ := GetTags(file)
	return tags
}
