package config

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type config struct {
	DB_Connection    string
	DB_MigrationFile string
}

var globalConfig *config

func GetConfig() *config {
	// lazy load our config singleton
	if globalConfig == nil {
		globalConfig = newHardCodedConfig()
		return globalConfig
	}

	return globalConfig
}

// create a new config object using environmental variables
func newEnvConfig() *config {
	panic("NOT YET IMPLEMENTED")
}

// create a new config from a file
func newConfigFromFile(fileName string) *config {
	panic("NOT YET IMPLEMENTED")
}

// only used for testing purposes, should likely be deleted before turn in
func newHardCodedConfig() *config {
	return &config{
		DB_Connection:    "postgres://postgres:root@localhost:5432/postgres?sslmode=disable",
		DB_MigrationFile: "file://migrations",
	}
}

// root@localhost:3306
