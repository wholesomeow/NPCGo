package texttypes

import "strings"

// A noun refers to a person, animal or thing.
type NounType struct {
	Noun string
}

type NounPhrase struct {
	Determiner *DeterminerType
	Adjective  *AdjectivePhrase
	Modifier   *NounPhrase
	Noun       NounType
}

// Builds an Noun Phrase struct from passed keywords.
// Takes in keywords, ID's the lexical category/POS of keywords
// and assigns them to a phraseType
func (phrase NounPhrase) BuildNounPhrase(noun NounType) NounPhrase {
	phrase.Noun = noun
	// NOTE(wholesomeow): Maybe have a "if proper noun/pronoun, just return else populate the rest of the phrase"?
	return phrase
}

// Gets all text from within a noun phrase
func GetNounText(phrase NounPhrase) string {
	text_slice := []string{}

	if phrase.Determiner != nil {
		determiner_text := phrase.Determiner.Determiner
		text_slice = append(text_slice, determiner_text)
	}
	if phrase.Adjective != nil {
		adjective_text := GetAdjectiveText(*phrase.Adjective)
		text_slice = append(text_slice, adjective_text)
	}
	if phrase.Modifier != nil {
		modifier_text := GetNounText(*phrase.Modifier)
		text_slice = append(text_slice, modifier_text)
	}
	noun_text := phrase.Noun.Noun
	text_slice = append(text_slice, noun_text)

	text := strings.Join(text_slice, " ")

	return text
}
