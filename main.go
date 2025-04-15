package main

import (
	"fmt"
	config "go/npcGen/configs"
	utilities "go/npcGen/internal/utilities"
	npcgen "go/npcGen/pkg/npcGen"
	"log"
	"time"
)

func main() {
	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Run all Pre-Flight checks
	err = utilities.DBPreFlight(config)
	if err != nil {
		log.Fatalf("failure in DBPreFlight: %v", err)
	}

	// Create NPC
	start_proc := time.Now()
	npc_object, err := npcgen.CreateNPC(config)
	if err != nil {
		log.Fatal(err)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)

	fmt.Println("----- OUTPUT -----")

	fmt.Println(npc_object.DataToJSON())
	// fmt.Println(npc_object.OCEAN.Text)
	log.Printf("npc created... elapsed time: %s", time.Duration.String(elapsed_proc))
}
