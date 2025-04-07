package npc

import (
	"context"
	"fmt"
	"go/npcGen/configuration"
	"go/npcGen/database"
	"go/npcGen/utilities"
	"log"
	"math/rand"
	"strings"

	"github.com/jackc/pgx/v4"
)

type MarkovChain struct {
	attempts         int
	keys             []string
	n_gram           map[string][]string
	vowels           []string
	accepted_bigrams []string
}

// TODO(wholesomeow): Use RogueBasin link to create more advanced Markov Chain
// LINK: http://www.roguebasin.com/index.php?title=Names_from_a_high_order_Markov_Process_and_a_simplified_Katz_back-off_scheme
func (mc *MarkovChain) BuildNGram(config *configuration.Config, max_attempts int) error {
	n_grams := [][]string{}
	compilation := map[string][]string{}

	//Get data from some place here, if no data then error
	if config.Server.Mode == "dev-csv" {
		path := fmt.Sprintf("%s/%s", config.Database.CSVPath, "Fantasy_Names_NGrams.csv")
		var err error
		n_grams, err = utilities.ReadCSV(path, false)
		if err != nil {
			return err
		}
	} else {
		log.Print("starting NGram data collection")
		// Create DB Object
		var db *pgx.Conn
		var err error
		db, err = database.ConnectDatabase(config)
		if err != nil {
			return err
		}

		defer db.Close(context.Background())

		// Query for required data to generate NPC
		err = db.QueryRow(context.Background(), "SELECT * FROM ngram_fantasy").Scan(&n_grams)
		if err != nil {
			return err
		}
	}

	// Split n_gram values into key and value slices
	for idx, val := range n_grams {
		if idx != 0 {
			mc.keys = append(mc.keys, val[0])
			options := strings.Split(val[1], ",")
			compilation[val[0]] = options
		}
	}

	mc.n_gram = compilation

	// TODO(wholesomeow): Move these to another config file
	vowles := []string{"a", "e", "i", "o", "u"}
	accepted_bigrams := []string{"br", "dr", "fr", "gr", "kr", "pr", "tr", "cr", "sn", "sw", "th",
		"sh", "ch", "cl", "sl", "sm", "sn", "sp", "st", "sk", "bl", "fl",
		"gl", "pl", "sl", "ll", "yl", "yv", "gh"}

	mc.attempts = max_attempts
	mc.vowels = vowles
	mc.accepted_bigrams = accepted_bigrams

	return nil
}

func (mchain *MarkovChain) GetStartPoint() (string, string) {
	keys := mchain.keys
	result := mchain.keys[rand.Intn(len(keys))]
	start_gram := result

	return result, start_gram
}

func (mchain *MarkovChain) MakeName() string {
	// TODO(wholesomeow): Implement better error checking and handling here
	log.Print("start of name creation")
	result, current_gram := mchain.GetStartPoint()

	for idx := range mchain.attempts {
		possibility := mchain.n_gram[current_gram]
		current_gram := possibility[rand.Intn(len(possibility))]

		if idx < mchain.attempts && current_gram == " " {
			break
		}
		result += current_gram
	}

	log.Print("end of name creation")
	return result
}

func (mchain *MarkovChain) CheckQuality(name string) bool {
	// Rules for name formatting are here
	log.Print("checking quality of name")
	if len(name) <= 3 {
		return false
	}

	for val := range len(name) - 1 {
		bigram := strings.ToLower(name[val : val+2])
		if utilities.SliceContainsString(string(bigram[0]), mchain.vowels) || utilities.SliceContainsString(string(bigram[1]), mchain.vowels) {
			continue
		} else if !utilities.SliceContainsString(string(bigram[0]), mchain.accepted_bigrams) {
			return false
		}
	}
	return true
}
