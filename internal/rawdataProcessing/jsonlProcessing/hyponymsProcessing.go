package rawdataproc

import "log"

type Hyponyms struct {
	Word   string
	Source string
	Dist   string
}

func ProcessHyponyms(data map[string]interface{}, categories []string, keys []string) []Hyponyms {
	var hyponyms []Hyponyms

	if items, ok := data[categories[0]].([]interface{}); ok {
		for _, item := range items {
			if entry, ok := item.(map[string]interface{}); ok {
				hyp := Hyponyms{
					Word:   JSONLStringParse(entry, keys[0]),
					Source: JSONLStringParse(entry, keys[1]),
					Dist:   JSONLStringParse(entry, keys[2]),
				}
				if hyp.Word != "" || hyp.Source != "" || hyp.Dist != "" {
					hyponyms = append(hyponyms, hyp)
				}
			} else {
				log.Printf("[ProcessHyponyms] - Issue with entry selection from category: %s", categories[0])
			}
		}
	}

	return hyponyms
}
