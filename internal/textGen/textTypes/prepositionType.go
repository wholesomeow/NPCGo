package texttypes

import "strings"

// Prepositions describe the relationship between words from the major word classes.
type PrepositionType struct {
	Preposition string
}

type PrepositionalPhrase struct {
	Verb        VerbType
	Preposition *PrepositionType
}

// Gets all text from within a preposition phrase
func GetPropositionText(phrase PrepositionalPhrase) string {
	text_slice := []string{}

	// Get text from Verb struct
	verb_text := phrase.Verb.Verb
	text_slice = append(text_slice, verb_text)
	if phrase.Preposition != nil {
		preposition_text := phrase.Preposition.Preposition
		text_slice = append(text_slice, preposition_text)
	}

	text := strings.Join(text_slice, " ")

	return text
}
