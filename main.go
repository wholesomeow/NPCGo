package main

import (
	"fmt"
	"go/npcGen/configuration"
	"go/npcGen/database"
	"go/npcGen/npc"
	"go/npcGen/utilities"
	"log"
	"os"
	"strings"
)

func dbPreFlight(config *configuration.Config) {
	log.Print("starting database pre-flight checks")

	// Check all required files in database/rawdata exist
	csv_path := config.Database.CSVPath
	log.Printf("Check for required files at %s", csv_path)
	for _, file := range config.Database.RequiredFiles {
		required_path := fmt.Sprintf("%s/%s", csv_path, file)
		utilities.CheckFilePath(required_path, true)
	}

	// Check all optional files in database/rawdata exist
	log.Printf("Check for optional files at %s", csv_path)
	optional_found := []bool{}
	for _, file := range config.Database.OptionalFiles {
		optional_path := fmt.Sprintf("%s/%s", csv_path, file)
		found := utilities.CheckFilePath(optional_path, false)
		optional_found = append(optional_found, found)

		// TODO(wholesomeow): If file(s) not found, ask what should be rebuilt
		// TODO(wholesomeow): Add force rebuild to CLI portion
	}

	// Build Optional data if files don't exist
	// TODO(wholesomeow): Add the rest of the optional files data processing - csv and json
	if !optional_found[0] {
		database.BuildNGramFromData(config)
	}
}

func main() {
	// Usage: go run main.go ./configuration/dbconf.yaml UP
	// Read in Database Config file
	var config configuration.Config
	conf_path := os.Args[1]
	log.Printf("database conf file at path %s", conf_path)
	utilities.ReadConfig(conf_path, &config)

	// Run all Pre-Flight checks
	dbPreFlight(&config)

	// If SKIP, then skip all database portions and build datafiles normally
	if os.Args[2] != "SKIP" {
		// Create and populate Database if not already done
		database.ConectDatabase(&config)
		// TODO(wholesomeow): Figure out why program is exiting after migrations complete
		arg := os.Args[2]
		option := strings.ToUpper(arg)
		database.MigrateDB(&config, option)
	} else {
		log.Printf("reading in command line option %s", os.Args[2])
	}

	// Create NPC
	npc_object := npc.CreateNPC(&config)
	fmt.Printf("npc_object: %v\n", npc_object)
}
