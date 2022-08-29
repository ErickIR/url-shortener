package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}

func GetDuration(key string, defaultValue time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return time.Duration(i) * time.Second
}

func GetBool(varName string, defaultValue bool) bool {
	val, ok := os.LookupEnv(varName)
	if !ok {
		return defaultValue
	}

	switch val {
	case "1", "t", "T", "true", "TRUE", "True", "yes", "Yes", "YES":
		return true
	case "0", "f", "F", "false", "FALSE", "False", "no", "No", "NO":
		return false
	}

	return defaultValue
}
