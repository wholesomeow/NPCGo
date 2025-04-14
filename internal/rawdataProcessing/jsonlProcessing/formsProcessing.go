package rawdataproc

import "log"

type Forms struct {
	Form     string
	FormsTag []string
}

func ProcessForms(data map[string]interface{}, categories []string, keys []string) []Forms {
	var forms []Forms

	if items, ok := data[categories[0]].([]interface{}); ok {
		for _, item := range items {
			if entry, ok := item.(map[string]interface{}); ok {
				frm := Forms{
					Form:     JSONLStringParse(entry, keys[0]),
					FormsTag: JSONLSliceParse(entry, categories[0], keys[1]),
				}
				if frm.Form != "" || len(frm.FormsTag) > 0 {
					forms = append(forms, frm)
				}
			} else {
				log.Printf("[ProcessForms] - Issue with entry selection from category: %s", categories[0])
			}
		}
	}

	return forms
}
