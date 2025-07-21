package utilities

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	config "github.com/wholesomeow/npcGo/configs"
	utilities "github.com/wholesomeow/npcGo/internal/utilities"
)

func ConnectDatabase(config *config.Config) (*sql.DB, error) {
	dbPath := config.Database.Path

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create db directory: %w", err)
	}

	log.Printf("Connecting to SQLite database at %s", dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping sqlite database: %w", err)
	}

	log.Printf("Successfully connected to SQLite database at %s", dbPath)
	return db, nil
}

func join(list []string, sep string) string {
	out := ""
	for i, item := range list {
		if i > 0 {
			out += sep
		}
		out += item
	}
	return out
}

func seedDatabase(db *sql.DB, config *config.Config) error {
	csvPath := config.Database.CSVPath
	jsonPath := config.Database.JSONPath

	// Read in all the files from the config
	files := config.Database.Files
	for file := range files {
		log.Printf("Reading in ")
		ext := strings.ToLower(filepath.Ext(files[file].Filename))

		if ext == ".csv" {
			filePath := csvPath + files[file].Filename
			data, err := utilities.ReadCSV(filePath, false)
			if err != nil {
				return err
			}
			header := data[0]

			// Seed csv data into database; First row is the header
			for _, row := range data[1:] {
				if len(row) != len(header) {
					return fmt.Errorf("column mismatch in row: %v", row)
				}

				// Create SQL INSERT
				qMarks := make([]string, len(header))
				for i := range header {
					qMarks[i] = "?"
				}

				query := fmt.Sprintf(
					"INSERT INTO %s (%s) VALUES (%s);",
					files[file].Tablename,
					join(header, ","),
					join(qMarks, ","),
				)

				vals := make([]interface{}, len(row))
				for i, v := range row {
					vals[i] = v
				}

				if _, err := db.Exec(query, vals...); err != nil {
					return fmt.Errorf("insert failed: %w", err)
				}
			}
		} else if ext == ".json" {
			filePath := jsonPath + files[file].Filename
			data, err := utilities.ReadJSON(filePath)
			if err != nil {
				return err
			}

			// Seed json data into database
			decoder := json.NewDecoder(bytes.NewReader(data))
			decoder.UseNumber() // Preserve number formatting for ints/floats

			var rows []map[string]interface{}
			if err := decoder.Decode(&rows); err != nil {
				return fmt.Errorf("failed to parse JSON: %w", err)
			}

			if len(rows) == 0 {
				return fmt.Errorf("no data to insert")
			}

			// Extract column names from the first row
			var columns []string
			for col := range rows[0] {
				columns = append(columns, col)
			}

			// Prepare insert statement
			columnList := strings.Join(columns, ", ")
			placeholderList := strings.Repeat("?, ", len(columns))
			placeholderList = strings.TrimSuffix(placeholderList, ", ")

			query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`,
				files[file].Tablename,
				columnList,
				placeholderList,
			)

			stmt, err := db.Prepare(query)
			if err != nil {
				return fmt.Errorf("failed to prepare insert statement: %w", err)
			}
			defer stmt.Close()

			// Insert each row
			for _, row := range rows {
				var values []interface{}
				for _, col := range columns {
					values = append(values, row[col])
				}
				_, err := stmt.Exec(values...)
				if err != nil {
					log.Printf("insert failed for row: %+v", row)
					return fmt.Errorf("insert error: %w", err)
				}
			}

			log.Printf("successfully inserted %d rows into %s", len(rows), files[file].Tablename)
		}
	}

	return nil
}

func InitDatabase(config *config.Config) error {
	dbPath := config.Database.Path

	// Ensure db directory exists
	if err := os.MkdirAll("db", 0755); err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := seedDatabase(db, config); err != nil {
		return err
	}

	fmt.Println("Database initialized successfully.")

	return nil
}
