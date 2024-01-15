package general

import (
	"os"
)

// GetEnv reads from environment and returns a default value as a string
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
