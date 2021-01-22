package env

import (
	"fmt"
	"os"
)

func Get(name string) (string, error) {
	v, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("missing env variable: %s", name)
	}
	return v, nil
}

func GetWithFallback(name string, fallback string) string {
	v := os.Getenv(name)
	if len(v) == 0 {
		return fallback
	}
	return v
}
