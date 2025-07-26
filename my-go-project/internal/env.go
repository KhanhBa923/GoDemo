package internal

import (
	"os"
	"strconv"
)

func GetStringEnv(key string, fallback string) string {
	os.LookupEnv(key)
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
func GetIntEnv(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
