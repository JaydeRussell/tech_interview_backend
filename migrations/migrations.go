package migrations

import (
	"log"

	"github.com/JaydeRussell/tech_interview_backend/config"
	"github.com/golang-migrate/migrate/v4"
)

// location of the sql migration files.
// 'file://*' specifies that we're working with a local file system
const migrationFolder = "file://migrations"

// runs all of the *.up.sql files in the migrations folder
func RunMigrations() {
	c := config.GetConfig()

	m, err := migrate.New(
		migrationFolder,
		c.DB_Connection,
	)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

// runs all of the *.down.sql files in the migrations folder, designed to clean up the migrations
func MigrationsDown() {
	c := config.GetConfig()

	m, err := migrate.New(
		migrationFolder,
		c.DB_Connection,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
