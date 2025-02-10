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

// TODO(wholesomeow): Properly implement this to check for required files from config and rebuild optional files if not found
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
	utilities.ReadConfig(conf_path, &config)

	// Run all Pre-Flight checks
	dbPreFlight(&config)

	// Create and populate Database if not already done
	database.ConectDatabase(&config)
	// TODO(wholesomeow): Figure out why program is exiting after migrations complete
	arg := os.Args[2]
	option := strings.ToUpper(arg)
	database.MigrateDB(&config, option)

	// Create NPC
	var npc npc.NPCBase

	npc.UUID = 412341234
	npc.Name = "Test Name"
	npc.Social_Network = [3]string{"person_1", "person_2", "person_3"}
	npc.CS_Dimension = "up"
	npc.REI_Data = "down"

	npc.DisplayName()
}
