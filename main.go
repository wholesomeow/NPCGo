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

func main() {
	// Read in Database Config file
	var config configuration.Config
	conf_path := "configuration/dbconf.yml"
	log.Printf("database conf file at path %s", conf_path)
	utilities.ReadConfig(conf_path, &config)

	// Run all Pre-Flight checks
	err := database.DBPreFlight(&config)
	if err != nil {
		log.Fatalf("failure in DBPreFlight: %s ", err)
	}

	// Initialize database per server mode selection
	mode := strings.ToLower(config.Server.Mode)
	log.Printf("reading in config mode option %s", mode)
	switch mode {
	case "dev-db":
		// Create and populate Database if not already done
		// conn, _ := database.ConectDatabase(&config)
		// database.MigrateDB(&config, conn, "UP")
		err := database.InitDB(&config)
		if err != nil {
			log.Fatal("failed to init database")
		}
	case "dev-csv":
		log.Print("Skipping database initialization")
	default:
		log.Fatalf("no mode matching %s. Please check configurations/dcbonf.yaml", mode)
	}

	// Create NPC
	npc_object, err := npc.CreateNPC(&config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----- OUTPUT -----")
	fmt.Println(npc.DataToJSON(npc_object))
	// fmt.Println(npc_object.OCEAN.Text)
}
