package utilities

import (
	"fmt"
	"log"
	"strings"

	config "github.com/wholesomeow/npcGo/configs"
)

type FoundData struct {
	Filename string
	Header   bool
}

func DBPreFlight(config *config.Config) error {
	log.Print("starting database pre-flight checks")

	// Check all required files in database/rawdata exist
	csv_path := config.Database.CSVPath
	json_path := config.Database.JSONPath
	found := []FoundData{}
	var path string

	log.Print("Check for required files")
	for _, file := range config.Database.Files {
		split := strings.Split(file.Filename, ".")
		suffix := split[len(split)-1]

		if suffix == "csv" {
			path = fmt.Sprintf("%s/%s", csv_path, file.Filename)
		} else if suffix == "json" {
			path = fmt.Sprintf("%s/%s", json_path, file.Filename)
		}

		if !CheckFilePath(path, file.Header) && file.Required {
			found_data := FoundData{Filename: file.Filename, Header: file.Header}
			found = append(found, found_data)
		}
	}

	// Build Optional data if files don't exist
	if len(found) >= 0 {
		var err error
		for _, val := range found {
			switch val.Filename {
			case "Fantasy_Names_NGrams.csv":
				err = BuildNGramFromData(config, val)
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
