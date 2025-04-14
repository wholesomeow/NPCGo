package rawdataproc

import "log"

type Etymology struct {
	Name      string
	Expansion string
}

func ProcessEtymology(data map[string]interface{}, categories []string, keys []string) []Etymology {
	var etymology []Etymology

	if items, ok := data[categories[0]].([]interface{}); ok {
		for _, item := range items {
			if entry, ok := item.(map[string]interface{}); ok {
				ety := Etymology{
					Name:      JSONLStringParse(entry, keys[0]),
					Expansion: JSONLStringParse(entry, keys[1]),
				}
				if ety.Name != "" || ety.Expansion != "" {
					etymology = append(etymology, ety)
				}
			} else {
				log.Printf("[ProcessEtymology] - Issue with entry selection from category: %s", categories[0])
			}
		}
	}

	return etymology
}
