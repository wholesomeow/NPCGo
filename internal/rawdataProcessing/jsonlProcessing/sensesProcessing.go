package rawdataproc

import (
	"log"
)

type Senses struct {
	Links    []string
	Glosses  []string
	Examples []string
}

func ProcessSenses(data map[string]interface{}, categories []string, keys []string) []Senses {
	var senses []Senses

	if items, ok := data[categories[0]].([]interface{}); ok {
		for _, item := range items {
			if entry, ok := item.(map[string]interface{}); ok {
				sen := Senses{
					Links:    JSONLSliceParse(entry, categories[1], keys[0]),
					Glosses:  JSONLSliceParse(entry, categories[2], keys[1]),
					Examples: JSONLSliceParse(entry, categories[3], keys[2]),
				}
				if len(sen.Links) > 0 || len(sen.Glosses) > 0 || len(sen.Examples) > 0 {
					senses = append(senses, sen)
				}
			} else {
				log.Printf("[ProcessSenses] - Issue with entry selection from category: %s", categories[0])
			}
		}
	}

	return senses
}
