package namegen

import (
	"log"
	"time"
)

func CreateName(version int) (string, error) {
	var name string
	max_attempts := 6
	var end_proc time.Time

	start_proc := time.Now()
	if version == 1 {
		var mchain MarkovChain

		// This queries the database for the precompiled n-gram data
		log.Print("starting ngram build")
		err := mchain.BuildNGram(max_attempts)
		if err != nil {
			return name, err
		}

		log.Print("starting name creation, using version 1")
		for count := range max_attempts {
			log.Printf("name creation attempt %d", count)
			name = mchain.MakeName()
			if mchain.CheckQuality(name) {
				break
			}
			log.Printf("name %s doesn't meet quality check... moving on to next attempt", name)
		}
		end_proc = time.Now()
	} else if version == 2 {
		log.Print("starting name creation, using version 2")

		names, err := QueryNames()
		if err != nil {
			return name, err
		}

		mchain := MakeGenerator(names, 3, .001)
		log.Print("start of name creation")

		name = mchain.Generate()

		log.Print("end of name creation")

		end_proc = time.Now()
	}

	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("name creation completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	return name, nil
}
