package texttypes

import "strings"

// A verb refers to an action, event or state.
type VerbType struct {
	Verb        string
	Object      *NounPhrase
	Preposition *PrepositionalPhrase
	Clause      *ClauseType
}

type VerbPhrase struct {
	Verb   VerbType
	Adverb *AdverbPhrase
}

// Builds an Verb Phrase struct from passed keywords.
// Takes in keywords, ID's the lexical category/POS of keywords
// and assigns them to a phraseType
func (phrase VerbPhrase) BuildVerbPhrase(verb VerbType) VerbPhrase {
	phrase.Verb = verb
	return phrase
}

// Gets all text from within a verb phrase
func GetVerbText(phrase VerbPhrase) string {
	text_slice := []string{}

	// Get text from Verb struct
	verb_text := phrase.Verb.Verb
	text_slice = append(text_slice, verb_text)
	if phrase.Verb.Object != nil {
		object_text := GetNounText(*phrase.Verb.Object)
		text_slice = append(text_slice, object_text)
	}
	if phrase.Verb.Preposition != nil {
		preposition_text := GetPropositionText(*phrase.Verb.Preposition)
		text_slice = append(text_slice, preposition_text)
	}
	if phrase.Verb.Clause != nil {
		clause_text := GetClauseText(*phrase.Verb.Clause)
		text_slice = append(text_slice, clause_text)
	}

	// Get text from Verb phrase struct
	if phrase.Adverb != nil {
		adverb_text := GetAdverbText(*phrase.Adverb)
		text_slice = append(text_slice, adverb_text)
	}

	text := strings.Join(text_slice, " ")

	return text
}
