package namegen

import (
	"context"
	"log"
	"math/rand"
	"strings"

	config "github.com/wholesomeow/npcGo/configs"
	utilities "github.com/wholesomeow/npcGo/internal/utilities"

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
func (mc *MarkovChain) BuildNGram(config *config.Config, max_attempts int) error {
	n_grams := [][]string{}
	compilation := map[string][]string{}

	log.Print("starting NGram data collection")
	// Create DB Object
	var db *pgx.Conn
	var err error
	db, err = utilities.ConnectDatabase(config)
	if err != nil {
		return err
	}

	defer db.Close(context.Background())

	// Query for required data to generate NPC
	var rows pgx.Rows
	log.Print("querying db for ngram data")
	rows, err = db.Query(context.Background(), "SELECT * FROM generator.ngram_fantasy")
	if err != nil {
		return err
	}

	defer rows.Close()

	// Iterate through query result
	log.Print("marshalling query data to slice")
	for rows.Next() {
		var ngram_id int
		var ngram_value string
		var ngram_posibility string
		var tmp_gram []string

		err := rows.Scan(&ngram_id, &ngram_value, &ngram_posibility)
		if err != nil {
			return err
		}

		tmp_gram = append(tmp_gram, ngram_value)
		tmp_gram = append(tmp_gram, ngram_posibility)

		n_grams = append(n_grams, tmp_gram)
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
	log.Printf("checking quality of name: %s", name)
	if len(name) <= 3 {
		return false
	}

	// Rules for name formatting are here
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
