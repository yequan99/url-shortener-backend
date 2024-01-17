package general

import (
	"fmt"
	"os"
)

// GetEnv reads from environment and returns a default value as a string
func GetJWTSecret(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if exists {
		return value, nil
	}

	return value, fmt.Errorf("No JWT Secret in env")
}
