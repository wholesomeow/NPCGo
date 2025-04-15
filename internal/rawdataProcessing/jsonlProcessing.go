package rawdataproc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Etymology struct {
	Name      string
	Expansion string
}

type Forms struct {
	Form     string
	FormsTag []string
}

type Synonyms struct {
	Word string
	Tags []string
	Dist string
}

type Hyponyms struct {
	Word   string
	Source string
	Dist   string
}

type Senses struct {
	AltOf    []string
	Links    []string
	Glosses  []string
	Examples []string
}

type Entry struct {
	Word            string
	POS             string
	Language        string
	LangCode        string
	Antonyms        []string
	Hypernyms       []string
	Related         []string
	Hyphenation     []string
	Derived         []string
	CoordinateTerms []string
	EtymologyList   []Etymology
	FormsList       []Forms
	SynonymsList    []Synonyms
	HyponymsList    []Hyponyms
	SensesList      []Senses
}

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
	var entries []Entry
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
	sen_keys := []string{"alt_of", "links", "glosses", "text"}
	sen_categories := []string{"senses", "links", "glosses", "examples"}

	for data := range lineChannel {
		// Process Etymology per Entry
		if items, ok := data[ety_categories[0]].([]interface{}); ok {
			for _, item := range items {
				if entry, ok := item.(map[string]interface{}); ok {
					ety := Etymology{
						Name:      JSONLStringParse(entry, ety_keys[0]),
						Expansion: JSONLStringParse(entry, ety_keys[1]),
					}
					if ety.Name != "" || ety.Expansion != "" {
						etymology = append(etymology, ety)
					}
				} else {
					log.Printf("[ProcessEtymology] - Issue with entry selection from category: %s", ety_categories[0])
				}
			}
		}

		// Process Forms per Entry
		if items, ok := data[frm_categories[0]].([]interface{}); ok {
			for _, item := range items {
				if entry, ok := item.(map[string]interface{}); ok {
					frm := Forms{
						Form:     JSONLStringParse(entry, frm_keys[0]),
						FormsTag: JSONLSliceParse(entry, frm_categories[0], frm_keys[1]),
					}
					if frm.Form != "" || len(frm.FormsTag) > 0 {
						forms = append(forms, frm)
					}
				} else {
					log.Printf("[ProcessForms] - Issue with entry selection from category: %s", frm_categories[0])
				}
			}
		}

		// Process Synonyms per Entry
		if items, ok := data[syn_categories[0]].([]interface{}); ok {
			for _, item := range items {
				if entry, ok := item.(map[string]interface{}); ok {
					syn := Synonyms{
						Word: JSONLStringParse(entry, syn_keys[0]),
						Tags: JSONLSliceParse(entry, syn_categories[0], syn_keys[1]),
						Dist: JSONLStringParse(entry, syn_keys[2]),
					}
					if syn.Word != "" || len(syn.Tags) > 0 || syn.Dist != "" {
						synonyms = append(synonyms, syn)
					}
				} else {
					log.Printf("[ProcessSynonyms] - Issue with entry selection from category: %s", syn_categories[0])
				}
			}
		}

		// Process Hyponyms per Entry
		if items, ok := data[hpy_categories[0]].([]interface{}); ok {
			for _, item := range items {
				if entry, ok := item.(map[string]interface{}); ok {
					hyp := Hyponyms{
						Word:   JSONLStringParse(entry, hyp_keys[0]),
						Source: JSONLStringParse(entry, hyp_keys[1]),
						Dist:   JSONLStringParse(entry, hyp_keys[2]),
					}
					if hyp.Word != "" || hyp.Source != "" || hyp.Dist != "" {
						hyponyms = append(hyponyms, hyp)
					}
				} else {
					log.Printf("[ProcessHyponyms] - Issue with entry selection from category: %s", hpy_categories[0])
				}
			}
		}

		// Process Senses per Entry
		if items, ok := data[sen_categories[0]].([]interface{}); ok {
			for _, item := range items {
				if entry, ok := item.(map[string]interface{}); ok {
					sen := Senses{
						AltOf:    JSONLSliceParse(entry, sen_categories[0], sen_keys[0]),
						Links:    JSONLSliceParse(entry, sen_categories[1], sen_keys[1]),
						Glosses:  JSONLSliceParse(entry, sen_categories[2], sen_keys[2]),
						Examples: JSONLSliceParse(entry, sen_categories[3], sen_keys[3]),
					}
					if len(sen.Links) > 0 || len(sen.Glosses) > 0 || len(sen.Examples) > 0 {
						senses = append(senses, sen)
					}
				} else {
					log.Printf("[ProcessSenses] - Issue with entry selection from category: %s", sen_categories[0])
				}
			}
		}

		entry := Entry{
			Word:            JSONLStringParse(data, def_keys[0]),
			POS:             JSONLStringParse(data, def_keys[1]),
			Language:        JSONLStringParse(data, def_keys[2]),
			LangCode:        JSONLStringParse(data, def_keys[3]),
			Antonyms:        JSONLSliceParse(data, def_categories[0], def_keys[0]),
			Hypernyms:       JSONLSliceParse(data, def_categories[1], def_keys[0]),
			Related:         JSONLSliceParse(data, def_categories[2], def_keys[0]),
			Hyphenation:     JSONLSliceParse(data, def_categories[3], def_keys[0]),
			Derived:         JSONLSliceParse(data, def_categories[4], def_keys[0]),
			CoordinateTerms: JSONLSliceParse(data, def_categories[5], def_keys[0]),
			EtymologyList:   etymology,
			FormsList:       forms,
			SynonymsList:    synonyms,
			HyponymsList:    hyponyms,
			SensesList:      senses,
		}
		entries = append(entries, entry)
	}
	<-done

	log.Printf("----- %s results -----", filepath)
	log.Printf("Entries length: %d", len(entries))
	log.Printf("Etymology length: %d", len(etymology))
	log.Printf("Forms length: %d", len(forms))
	log.Printf("Synoyms length: %d", len(synonyms))
	log.Printf("Hyponyms length: %d", len(hyponyms))
	log.Printf("Senses length: %d", len(senses))

	return nil
}

// ExtractFirstJSONL reads a JSONL file and writes the first entry to a JSON file
func ExtractFirstJSONL(inputPath string, outputPath string) error {
	// Open the JSONL file
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Unmarshal to check if it's valid JSON
		var jsonObj interface{}
		err := json.Unmarshal([]byte(line), &jsonObj)
		if err != nil {
			return fmt.Errorf("invalid JSON in line: %w", err)
		}

		// Marshal it with indentation for readability
		formattedJSON, err := json.MarshalIndent(jsonObj, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to format JSON: %w", err)
		}

		// Write to output file
		err = os.WriteFile(outputPath, formattedJSON, 0644)
		if err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}

		fmt.Printf("First JSON entry written to %s\n", outputPath)
		return nil
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return fmt.Errorf("no valid JSON entries found in file")
}
