package rawdataproc

type Definition struct {
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
}

func ProcessDefinition(data map[string]interface{}, categories []string, keys []string) Definition {
	def := Definition{
		Word:            JSONLStringParse(data, keys[0]),
		POS:             JSONLStringParse(data, keys[1]),
		Language:        JSONLStringParse(data, keys[2]),
		LangCode:        JSONLStringParse(data, keys[3]),
		Antonyms:        JSONLSliceParse(data, categories[0], keys[0]),
		Hypernyms:       JSONLSliceParse(data, categories[1], keys[0]),
		Related:         JSONLSliceParse(data, categories[2], keys[0]),
		Hyphenation:     JSONLSliceParse(data, categories[3], keys[0]),
		Derived:         JSONLSliceParse(data, categories[4], keys[0]),
		CoordinateTerms: JSONLSliceParse(data, categories[5], keys[0]),
	}

	return def
}
