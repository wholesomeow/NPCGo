package main

import (
	"fmt"
	"go/npcGen/configuration"
	"go/npcGen/database"
	"go/npcGen/npc"
	"log"
)

func main() {
	// Read in Database Config file
	config, err := configuration.ReadConfig("configuration/dbconf.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Run all Pre-Flight checks
	err = database.DBPreFlight(config)
	if err != nil {
		log.Fatalf("failure in DBPreFlight: %s ", err)
	}

	// Create NPC
	npc_object, err := npc.CreateNPC(config)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print("testing")

	fmt.Println("----- OUTPUT -----")
	fmt.Println(npc_object.DataToJSON())
	// fmt.Println(npc_object.OCEAN.Text)
}
