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

func GetConfig() *config {
	// lazy load our config singleton
	if globalConfig == nil {
		globalConfig = newEnvConfig()
		return globalConfig
	}

	return globalConfig
}

// create a new config object using environmental variables
func newEnvConfig() *config {
	return &config{
		DB_Connection: mustExist("DB_CONNECTION"),
	}
}

func mustExist(v string) string {
	possibleVariable := os.Getenv(v)
	if possibleVariable == "" {
		log.Fatalf("could not find %s or was empty. Must be provided", v)
	}
	return possibleVariable
}
