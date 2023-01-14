package sql_migrations

import (
	"log"

	"github.com/JaydeRussell/tech_interview_backend/config"
	"github.com/golang-migrate/migrate/v4"
)

func RunMigrations() {
	c := config.GetConfig()

	m, err := migrate.New(
		c.SQL_MigrationFile,
		c.SQL_Connection,
	)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal()
	}
}
