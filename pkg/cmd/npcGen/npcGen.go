package main

import (
	"fmt"
	config "go/npcGen/configs"
	rawdataproc "go/npcGen/internal/rawdataProcessing/jsonlProcessing"
	utilities "go/npcGen/internal/utilities"
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

	// Process JSONL Files
	start_proc := time.Now()
	log.Print("starting JSON processing")
	err = rawdataproc.ProcessJSONL("database/rawdata/jsonl/pos-adj.jsonl")
	if err != nil {
		log.Fatalf("failure in JSONL data processing: %v", err)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("processing completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	// Create NPC
	// npc_object, err := npc.CreateNPC(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("----- OUTPUT -----")

	// fmt.Println(npc_object.DataToJSON())
	// fmt.Println(npc_object.OCEAN.Text)
}
