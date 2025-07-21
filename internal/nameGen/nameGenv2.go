package namegen

import (
	"log"
	"strings"

	config "github.com/wholesomeow/npcGo/configs"
	db "github.com/wholesomeow/npcGo/db"
	"github.com/wholesomeow/npcGo/internal/utilities"
)

// Struct to contain each character in the supporting corpus,
// here called support, and it's amount of appearances.
// Used to determine the probability of a charcter by
// dividing the count of a specific characer (rune) by
// the total observed counts of that character
type Catagorial struct {
	Counts map[rune]float64 // Map of appearances of each character
	Total  float64          // Total sum of appearances
}

type MarkovChainV2 struct {
	Support  []rune
	Order    int
	Prior    float64
	Boundary rune
	Prefix   []rune
	Counts   map[string]*Catagorial
}

func QueryNames() ([]string, error) {
	log.Print("starting NGram data collection")
	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		return nil, err
	}

	// Create DB Object
	database, err := db.ConnectDatabase(config)
	if err != nil {
		return nil, err
	}

	defer database.Close()

	// Query for required data to generate NPC
	log.Print("querying db for names data")
	rows, err := database.Query("SELECT * FROM generator.names_fantasy")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatalf("Row scan failed: %v\n", err)
		}
		results = append(results, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Row iteration error: %v\n", err)
	}

	return results, nil
}

// Initializes a categorical struct with each rune in a
// provided supporting corpus with it's initial appearances
// set to the prior, to smooth out a count when there would
// otherwise not be one
func MakeCategorial(support []rune, prior float64) *Catagorial {
	counts := make(map[rune]float64)

	// Assigns the prior weight to each character in the support corpus
	for _, ch := range support {
		counts[ch] = prior
	}

	//
	return &Catagorial{
		Counts: counts,
		Total:  float64(len(support)) * prior,
	}
}

// Observes each event (character) and adds the amount of times
// that event has been observed to the categorical struct
func (categorical *Catagorial) Observe(event rune, count float64) {
	categorical.Counts[event] += count
	categorical.Total += count
}

// Selects a character based off a random float sampled within
// range of the total observed counts of an event
func (categorical *Catagorial) Sample() rune {
	sample := utilities.RandomRange64(0.0, categorical.Total)

	for event, count := range categorical.Counts {
		if sample <= count {
			return event
		}
		sample -= count
	}

	// Fallback if sample is never met, just return
	// the first event in Counts
	for event := range categorical.Counts {
		return event
	}

	// If all else fails, return a "fail character"
	return '?'
}

// Function to call when iterating through the categorical.
// Not to be confused with iterating through Counts
func (categorical *Catagorial) Get(event rune) float64 {
	return categorical.Counts[event] / categorical.Total
}

func MakeMarkovChain(support []rune, order int, prior float64, boundary rune) *MarkovChainV2 {
	// Generate prefix
	prefix := make([]rune, order)
	for i := range prefix {
		prefix[i] = boundary
	}

	return &MarkovChainV2{
		Support:  support,
		Order:    order,
		Prior:    prior,
		Boundary: boundary,
		Prefix:   prefix,
		Counts:   make(map[string]*Catagorial),
	}
}

func (markovchain *MarkovChainV2) getCategorial(context []rune) *Catagorial {
	context_key := string(context)
	if _, ok := markovchain.Counts[context_key]; !ok {
		markovchain.Counts[context_key] = MakeCategorial(markovchain.Support, markovchain.Prior)
	}

	return markovchain.Counts[context_key]
}

func (markovchain *MarkovChainV2) backoff(context []rune) []rune {
	if len(context) > markovchain.Order {
		context = context[len(context)-markovchain.Order:]
	} else if len(context) < markovchain.Order {
		padding := make([]rune, markovchain.Order-len(context))

		for i := range padding {
			padding[i] = markovchain.Boundary
		}

		context = append(padding, context...)
	}

	for len(context) > 0 {
		context_key := context
		if _, ok := markovchain.Counts[string(context_key)]; ok {
			return context_key
		}

		context = context[1:]
	}

	return []rune{' '}
}

func (markovchain *MarkovChainV2) Observe(sequence string, count float64) {
	// TODO(wholesomeow): Implement prefix and suffix

	// Just calling the current value event as it's the
	// current event we will want to observe
	for i, event := range sequence[markovchain.Order:] {

		// Context gets set to sliding window behind the current index
		// to provide previous state to chain
		context := sequence[i-markovchain.Order : i]
		for j := range len(context) + 1 {
			markovchain.getCategorial([]rune(context[j:])).Observe(event, count)
		}
	}
}

func (markovchain *MarkovChainV2) Sample(context []rune) rune {
	context = markovchain.backoff(context)
	return markovchain.getCategorial([]rune(context)).Sample()
}

func (markovchain *MarkovChainV2) Generate() string {
	sequence := []rune{markovchain.Sample(markovchain.Prefix)}
	for len(sequence) > 0 {
		if sequence[len(sequence)-1] == markovchain.Boundary {
			break
		}

		sequence = append(sequence, markovchain.Sample(sequence))
	}

	return string(sequence[:len(sequence)-1])
}

func MakeGenerator(names_slice []string, order int, prior float64) *MarkovChainV2 {
	pre_support := make(map[rune]bool)
	pre_names := make(map[string]bool)

	for _, name := range names_slice {
		name = strings.TrimSpace(name)

		pre_names[name] = true // prevents duplicates

		// Converts strings to runes for the rest of the Markov Chain
		for _, char := range name {
			pre_support[char] = true // prevents duplicates
		}
	}

	boundary := '$'
	pre_support[boundary] = true

	// Takes the "immutable" support and coverts it into a regular slice
	support := []rune{}
	for char := range pre_support {
		support = append(support, char)
	}

	// Same here
	names := []string{}
	for name := range pre_names {
		names = append(names, name)
	}

	// Create the MarkovChainV2 model
	markovchain := MakeMarkovChain(support, order, prior, boundary)
	for _, name := range names {
		markovchain.Observe(name, 1.0)
	}

	return markovchain
}
