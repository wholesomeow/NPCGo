package database

import (
	"database/sql"
	"fmt"
	"go/npcGen/configuration"
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

func ConectDatabase(config *configuration.Config) {
	// Connect to database
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Database.Username,
		config.Database.Password,
		config.Server.Host,
		config.Server.Port,
		config.Database.DBName,
		config.Database.SSLMode,
	)
	database, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("validated connection arguments... opening connection to database now")

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Printf("successfully connected to database %s as user %s",
		config.Database.DBName, config.Database.Username)

	defer database.Close()
}

func MigrateDB(config *configuration.Config, arg string) {
	// Read config and start migration
	// TODO(wholesomeow): Research search_path settings
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=public",
		config.Database.Username,
		config.Database.Password,
		config.Server.Host,
		config.Server.Port,
		config.Database.DBName,
		config.Database.SSLMode,
	)
	migration_path := config.Database.MigrationPath
	m, err := migrate.New(migration_path, connStr)

	// TODO(wholesomeow): Better error handling here
	if err != nil {
		log.Fatal(err)
	}

	switch arg {
	case "UP":
		log.Printf("begining migration '%s' in database %s for user %s",
			arg, config.Database.DBName, config.Database.Username)
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "DOWN":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("no known argument provided")
	}
}
