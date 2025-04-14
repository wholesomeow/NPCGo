package rawdataproc

import "log"

type Synonyms struct {
	Word string
	Tags []string
	Dist string
}

func ProcessSynonyms(data map[string]interface{}, categories []string, keys []string) []Synonyms {
	var synonyms []Synonyms

	if items, ok := data[categories[0]].([]interface{}); ok {
		for _, item := range items {
			if entry, ok := item.(map[string]interface{}); ok {
				syn := Synonyms{
					Word: JSONLStringParse(entry, keys[0]),
					Tags: JSONLSliceParse(entry, categories[0], keys[1]),
					Dist: JSONLStringParse(entry, keys[2]),
				}
				if syn.Word != "" || len(syn.Tags) > 0 || syn.Dist != "" {
					synonyms = append(synonyms, syn)
				}
			} else {
				log.Printf("[ProcessSynonyms] - Issue with entry selection from category: %s", categories[0])
			}
		}
	}

	return synonyms
}
