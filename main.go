package main

import (
	"fmt"
	"log"
	"time"

	config "github.com/wholesomeow/npcGo/configs"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func main() {
	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		log.Fatal(err)
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
