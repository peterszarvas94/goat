package content

func GetTitle(matter map[string]any, def string) string {
	if title, ok := matter["title"].(string); ok {
		return title
	}
	return def
}
