package main

import (
	"fmt"
	"go/npcGen/configuration"
	"go/npcGen/database"
	"go/npcGen/npc"
	"go/npcGen/utilities"
	"log"
	"strings"
)

func dbPreFlight(config *configuration.Config) {
	log.Print("starting database pre-flight checks")

	// Check all required files in database/rawdata exist
	// TODO(wholesomeow): This sucks lol
	csv_path := config.Database.CSVPath
	json_path := config.Database.JSONPath
	var required_path string
	log.Print("Check for required files")
	for _, file := range config.Database.RequiredFiles {
		split := strings.Split(file, ".")
		suffix := split[len(split)-1]
		if suffix == "csv" {
			required_path = fmt.Sprintf("%s/%s", csv_path, file)
		} else if suffix == "json" {
			required_path = fmt.Sprintf("%s/%s", json_path, file)
		}
		utilities.CheckFilePath(required_path, true)
	}

	// Check all optional files in database/rawdata exist
	optional_found := []bool{}
	var optional_path string
	log.Print("Check for optional files")
	for _, file := range config.Database.OptionalFiles {
		split := strings.Split(file, ".")
		suffix := split[len(split)-1]
		if suffix == "csv" {
			optional_path = fmt.Sprintf("%s/%s", csv_path, file)
		} else if suffix == "json" {
			optional_path = fmt.Sprintf("%s/%s", json_path, file)
		}
		found := utilities.CheckFilePath(optional_path, false)
		optional_found = append(optional_found, found)
	}

	// Build Optional data if files don't exist
	// TODO(wholesomeow): Add the rest of the optional files data processing - csv and json
	if !optional_found[0] {
		database.BuildNGramFromData(config)
	}
}

func main() {
	// Read in Database Config file
	var config configuration.Config
	conf_path := "configuration/dbconf.yml"
	log.Printf("database conf file at path %s", conf_path)
	utilities.ReadConfig(conf_path, &config)

	// Run all Pre-Flight checks
	dbPreFlight(&config)

	// Initialize database per server mode selection
	mode := strings.ToLower(config.Server.Mode)
	log.Printf("reading in config mode option %s", mode)
	switch mode {
	case "dev-db":
		// Create and populate Database if not already done
		database.ConectDatabase(&config)
		database.MigrateDB(&config, "UP")
	case "dev-csv":
		log.Print("Skipping database initialization")
	default:
		log.Fatalf("no mode matching %s. Please check configurations/dcbonf.yaml", mode)
	}

	// Create NPC
	npc_object := npc.CreateNPC(&config)

	fmt.Println("----- OUTPUT -----")
	fmt.Println(npc.DataToJSON(npc_object))
	// fmt.Println(npc_object.OCEAN.Text)
}
