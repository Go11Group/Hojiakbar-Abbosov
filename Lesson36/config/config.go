package config

import (
	"fmt"
	"os"
)

var (
	DBHost     = getEnv("DB_HOST", "localhost")
	DBPort     = getEnv("DB_PORT", "5432")
	DBUser     = getEnv("DB_USER", "postgres")
	DBPassword = getEnv("DB_PASSWORD", "root")
	DBName     = getEnv("DB_NAME", "leetcode")
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

func GetDBURI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, DBHost, DBPort, DBName)
}
