package database

import (
	"context"
	"fmt"
	"go/npcGen/configuration"
	"go/npcGen/utilities"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

func ConectDatabase(config *configuration.Config) (*pgx.Conn, error) {
	log.Printf("connecting to database on %s as %s on port %d",
		config.Database.Hostname,
		config.Database.Username,
		config.Database.Port,
	)

	// Set connection string
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Hostname,
		config.Database.Port,
		config.Database.DBName,
		config.Database.SSLMode,
	)

	// Open connection to database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return conn, err
	}

	log.Print("validated connection arguments... opening connection to database now")

	// Test connection is good
	if err = conn.Ping(context.Background()); err != nil {
		return conn, err
	}

	log.Printf("successfully connected to database %s as user %s",
		config.Database.DBName,
		config.Database.Username,
	)

	return conn, nil
}

func transferData(config *configuration.Config, conn *pgx.Conn, data [][]string, count int) error {
	// Set variables
	file := config.Database.Files[count]
	var tx_data [][]interface{}

	// Start a transaction
	log.Print("starting transaction")
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("conn.Begin failed... error: %s", err)
	}
	defer tx.Rollback(context.Background())

	// TODO(wholesomeow): Prep schema here with strings.ToUpper()

	// NOTE(wholesomeow): CopyFromRows takes in an [][]interface, so need to convert data
	log.Printf("converting data from file %s to interface", file.Filename)
	for idx, record := range data {
		log.Printf("writing row %d of %d from file %s",
			idx,
			len(record),
			file.Filename,
		)
		row := make([]interface{}, len(record))
		for i, val := range record {
			// TODO(wholesomeow): Implement type conversion here
			row[i] = val
		}
		tx_data = append(tx_data, row)
	}

	// Insert data into table
	log.Printf("copy data into table %s", file.Tablename)
	copyCount, err := tx.CopyFrom(
		context.Background(),
		pgx.Identifier{file.Tablename},
		file.Schema,
		pgx.CopyFromRows(tx_data),
	)
	if err != nil {
		return err
	}

	log.Print("commiting transaction")
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	// Close connection to database
	defer conn.Close(context.Background())

	log.Printf(
		"successfully imported %d rows into %s: ",
		copyCount,
		file.Tablename,
	)

	return nil
}

func readRawData(config *configuration.Config, count int) ([][]string, error) {
	log.Print("begin writing file to database")
	// Set variables for data transaction
	csv_path := config.Database.CSVPath
	json_path := config.Database.JSONPath
	file := config.Database.Files[count]
	output := [][]string{}

	// Check filename suffix
	split := strings.Split(file.Filename, ".")
	suffix := split[len(split)-1]

	log.Printf("copy data from %s to %s", file.Filename, file.Tablename)
	if suffix == "csv" {
		full_path := fmt.Sprintf("%s/%s", csv_path, file.Filename)
		output, err := utilities.ReadCSV(full_path, file.Header)

		return output, err
	} else if suffix == "json" {
		full_path := fmt.Sprintf("%s/%s", json_path, file.Filename)
		json_data, err := utilities.ReadJSON(full_path)
		if err != nil {
			return output, err
		}

		output, err := utilities.JSONToSlice(json_data)
		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func CreateTable(config *configuration.Config, table_name string) error {
	// Connect to database
	conn, err := ConectDatabase(config)
	if err != nil {
		return err
	}

	// Create tables from table names
	migration_path := config.Database.MigrationPath
	table_path := fmt.Sprintf(
		"%s/create_table_%s.sql",
		migration_path,
		table_name,
	)
	_, err = conn.Exec(context.Background(), table_path)
	if err != nil {
		return err
	}
	return nil
}

func InitDB(config *configuration.Config) error {
	log.Print("begin database content init")
	// Connect to database
	conn, err := ConectDatabase(config)
	if err != nil {
		return fmt.Errorf("func ConnectDatabase error: %s", err)
	}

	// Create tables if they don't exist
	for _, file := range config.Database.Files {
		var n int64
		query := fmt.Sprintf(
			"select 1 from information_schema.tables where table_name=%s",
			file.Tablename,
		)
		err := conn.QueryRow(context.Background(), query, "ID").Scan(&n)
		if err == pgx.ErrNoRows {
			err := CreateTable(config, file.Filename)
			if err != nil {
				return fmt.Errorf("func CreateTable error: %s", err)
			}
		} else if err != nil {
			return err
		}
	}

	// Commit file raw data to table
	for idx := range config.Database.Files {
		data, err := readRawData(
			config,
			idx,
		)
		if err != nil {
			return fmt.Errorf("func readRawData error: %s", err)
		}

		err = transferData(config, conn, data, idx)
		if err != nil {
			return fmt.Errorf("func transferData error: %s", err)
		}
	}

	return nil
}

func MigrateDB(config *configuration.Config, conn *pgx.Conn, arg string) error {
	// Read config and start migration
	// TODO(wholesomeow): Research search_path settings
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=public",
		config.Database.Username,
		config.Database.Password,
		config.Database.Hostname,
		config.Database.Port,
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

	return nil
}
