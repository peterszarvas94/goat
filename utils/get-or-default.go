package utils

func GetOrDefault(m map[string]string, key, defaultValue string) string {
	if v, exists := m[key]; exists {
		return v
	}
	return defaultValue
}
