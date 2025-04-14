package rawdataproc

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func JSONLStringParse(data map[string]interface{}, key string) string {
	var out_word string
	if subdata, ok := data[key].(string); ok {
		out_word = subdata
	}

	return out_word
}

func JSONLParse(data interface{}, key string) string {
	var out_word string

	return out_word
}

func JSONLSliceParse(data map[string]interface{}, category string, key string) []string {
	var out_slice []string
	switch val := data[category].(type) {
	case []interface{}:
		for _, item := range val {
			switch v := item.(type) {
			case map[string]interface{}:
				out := JSONLStringParse(v, key)
				out_slice = append(out_slice, out)
			case []interface{}:
				for _, subVal := range v {
					if str, ok := subVal.(string); ok {
						out_slice = append(out_slice, str)
					}
				}
			default:
				if str, ok := item.(string); ok {
					out_slice = append(out_slice, str)
				}
			}
		}
	case map[string]interface{}:
		out := JSONLStringParse(val, key)
		out_slice = append(out_slice, out)
	case []map[string]interface{}:
		for _, item := range val {
			out := JSONLStringParse(item, key)
			out_slice = append(out_slice, out)
		}
	}

	return out_slice
}

func ProcessJSONL(filepath string) error {
	var definitions []Definition
	var etymology []Etymology
	var forms []Forms
	var synonyms []Synonyms
	var hyponyms []Hyponyms
	var senses []Senses

	// Open the JSONL file
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("could not open file: %v", err)
	}
	defer file.Close()

	// Create a scanner with increased buffer size
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	// Make Channel to recieve lines
	lineChannel := make(chan map[string]interface{})
	done := make(chan struct{})

	// Loop through each line in the file
	go func() {
		for scanner.Scan() {
			line := scanner.Text()

			// Parse the JSON data in the line
			data := make(map[string]interface{})
			if err := json.Unmarshal([]byte(line), &data); err != nil {
				log.Printf("could not unmarshal JSON: %v", err)
			}

			// Push current unmarshal to line channel
			lineChannel <- data
		}
		close(lineChannel)
		done <- struct{}{}
	}()

	def_keys := []string{"word", "pos", "lang", "lang_code"}
	def_categories := []string{"antonyms", "hypernyms", "related", "hyphenation", "derived", "coordinate_terms"}
	ety_keys := []string{"name", "expansion"}
	ety_categories := []string{"etymology_templates"}
	frm_keys := []string{"form", "tags"}
	frm_categories := []string{"forms"}
	syn_keys := []string{"word", "tags", "_dis1"}
	syn_categories := []string{"synonyms"}
	hyp_keys := []string{"word", "source", "_dis1"}
	hpy_categories := []string{"hyponyms"}
	sen_keys := []string{"links", "glosses", "text"}
	sen_categories := []string{"senses", "links", "glosses", "examples"}

	for data := range lineChannel {
		def := ProcessDefinition(data, def_categories, def_keys)
		definitions = append(definitions, def)

		etys := ProcessEtymology(data, ety_categories, ety_keys)
		etymology = append(etymology, etys...)

		frms := ProcessForms(data, frm_categories, frm_keys)
		forms = append(forms, frms...)

		syns := ProcessSynonyms(data, syn_categories, syn_keys)
		synonyms = append(synonyms, syns...)

		hpys := ProcessHyponyms(data, hpy_categories, hyp_keys)
		hyponyms = append(hyponyms, hpys...)

		sens := ProcessSenses(data, sen_categories, sen_keys)
		senses = append(senses, sens...)
	}
	<-done

	log.Printf("----- %s results -----", filepath)
	log.Printf("Definitions length: %d", len(definitions))
	log.Printf("Etymology length: %d", len(etymology))
	log.Printf("Forms length: %d", len(forms))
	log.Printf("Synoyms length: %d", len(synonyms))
	log.Printf("Hyponyms length: %d", len(hyponyms))
	log.Printf("Senses length: %d", len(senses))

	return nil
}
