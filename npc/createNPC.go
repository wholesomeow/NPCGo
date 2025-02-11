package npc

import (
	"go/npcGen/configuration"
	"log"
	"time"
)

func CreateName(config *configuration.Config) string {
	var mchain MarkovChain
	var name string
	max_attempts := 6

	buildNGram(&mchain, config, max_attempts)

	log.Print("starting name creation")
	start_proc := time.Now()
	for count := range max_attempts {
		log.Printf("name creation attempt %d", count)
		name = makeName(&mchain)
		if checkQuality(&mchain, name) {
			break
		}
		log.Printf("name %s doesn't meet quality check... moving on to next attempt", name)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("name creation completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	return name
}

func CreateNPC(config *configuration.Config) NPCBase {
	var npc NPCBase
	npc.Name = CreateName(config)

	return npc
}
