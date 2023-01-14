package config

type config struct {
	SQL_Connection    string
	SQL_MigrationFile string
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
		SQL_Connection:    "",
		SQL_MigrationFile: "",
	}
}
