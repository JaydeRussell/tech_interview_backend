package config

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type config struct {
	DB_Connection string
}

var globalConfig *config

// connects to the docker postgres created in the docker-compose.yml. This is purely for testing purposes
// I would likely never push a real db connection string to a public github repo, that'd be silly
const localPostgresAddress = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func GetConfig() *config {
	// lazy load our config singleton
	if globalConfig == nil {
		globalConfig = newEnvConfig()
		return globalConfig
	}

	return globalConfig
}

// attempts to use the environment variables, if one isn't available, then it will use a provided fallback value
func newEnvConfig() *config {
	return &config{
		DB_Connection: getWithFallBack("DB_CONNECTION", localPostgresAddress),
	}
}

func mustExist(v string) string {
	possibleVariable := os.Getenv(v)
	if possibleVariable == "" {
		log.Fatalf("could not find %s or was empty. Must be provided", v)
	}
	return possibleVariable
}

func getWithFallBack(varName, fallback string) string {
	possibleVariable := os.Getenv(varName)
	if possibleVariable == "" {
		return fallback
	}
	return possibleVariable
}
