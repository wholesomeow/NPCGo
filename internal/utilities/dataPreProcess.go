package utilities

import (
	"fmt"
	config "go/npcGen/configs"
	"log"
	"strings"
	"time"
)

func BuildNGramFromData(config *config.Config, file FoundData) error {
	path := fmt.Sprintf(
		"%s/%s",
		config.Database.CSVPath,
		file.Filename,
	)
	pre_name_data, err := ReadCSV(path, file.Header)
	if err != nil {
		return nil
	}
	var collected_strings strings.Builder

	// Dump the csv data into a single slice
	log.Print("pre-processing ngram data from csv")
	for _, val := range pre_name_data {
		for _, ele := range val {
			collected_strings.WriteString(fmt.Sprintf("%v ", ele))
		}
	}

	input := collected_strings.String()

	n_grams := map[string][]string{}
	log.Print("processing ngram data")
	start_proc := time.Now()
	for idx, val := range input {
		// FIXME(wholesomeow): This is a bad hack, not sure why it returns index out of range at 64681 when length is 64682
		if idx == len(input)-2 {
			break
		}

		// Checks if val is not in n-grams
		//   if not, create new entry and populate
		//   else, place appropriatly

		// TODO(wholesomeow): Something about this takes ~6.5 seconds, gotta be a faster way to iterate through this
		key := string(val)
		if _, exists := n_grams[key]; !exists {
			n_grams[key] = []string{}
		} else {
			bi_gram := string([]rune(input)[idx+1])
			n_grams[key] = append(n_grams[key], bi_gram)
		}
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("processing completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	// Dump n-gram into new csv in /database/rawdata/csv
	output := [][]string{}
	log.Print("formatting ngram data")
	start_format := time.Now()
	for key, val := range n_grams {
		tmp := []string{}
		tmp = append(tmp, key)
		cat_val := strings.Join(val, ",")
		tmp = append(tmp, cat_val)
		output = append(output, tmp)
	}
	end_format := time.Now()
	elapsed_format := end_format.Sub(start_format)
	log.Printf("formatting completed... elapsed time: %s", time.Duration.String(elapsed_format))

	WriteCSV(
		config.Database.CSVPath,
		file.Filename,
		output,
	)

	return nil
}
