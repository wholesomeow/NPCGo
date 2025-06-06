package namegen

import (
	"log"
	"time"

	config "github.com/wholesomeow/npcGo/configs"
)

func CreateName() (string, error) {
	var mchain MarkovChain
	var name string
	max_attempts := 6

	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		log.Fatal(err)
	}

	// This queries the database for the precompiled n-gram data
	log.Print("starting ngram build")
	err = mchain.BuildNGram(config, max_attempts)
	if err != nil {
		return name, err
	}

	log.Print("starting name creation")
	start_proc := time.Now()
	for count := range max_attempts {
		log.Printf("name creation attempt %d", count)
		name = mchain.MakeName()
		if mchain.CheckQuality(name) {
			break
		}
		log.Printf("name %s doesn't meet quality check... moving on to next attempt", name)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("name creation completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	return name, nil
}
