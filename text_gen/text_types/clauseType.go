package texttypes

import "strings"

type SubjectComplimentType struct {
	NounPhrase *NounPhrase
	Adjective  *AdjectivePhrase
}

type ObjectComplimentType struct {
	NounPhrase      *NounPhrase
	Adjective       *AdjectivePhrase
	AdjectivePhrase *AdjectivePhrase
}

type PrepositionalComplimentType struct {
	Preposition *PrepositionType
	NounPhrase  *NounPhrase
}

type ComplimentType struct {
	SubjectCompliment       *SubjectComplimentType
	ObjectCompliment        *ObjectComplimentType
	PrepositionalCompliment *PrepositionalComplimentType
}

type ClauseType struct {
	Subject    NounPhrase
	Verb       VerbPhrase
	Object     *NounPhrase
	Compliment *ComplimentType
	Adjunct    *AdverbPhrase
}

func GetComplimentText(phrase ComplimentType) string {
	text_slice := []string{}

	// Get Text from Subject Compliment struct
	if phrase.SubjectCompliment != nil {
		if phrase.SubjectCompliment.NounPhrase != nil {
			noun_text := GetNounText(*phrase.SubjectCompliment.NounPhrase)
			text_slice = append(text_slice, noun_text)
		}
		if phrase.SubjectCompliment.Adjective != nil {
			adjective_text := GetAdjectiveText(*phrase.SubjectCompliment.Adjective)
			text_slice = append(text_slice, adjective_text)
		}
	}

	// Get Text from Object Compliment struct
	if phrase.ObjectCompliment != nil {
		if phrase.ObjectCompliment.NounPhrase != nil {
			noun_text := GetNounText(*phrase.ObjectCompliment.NounPhrase)
			text_slice = append(text_slice, noun_text)
		}
		if phrase.ObjectCompliment.Adjective != nil {
			adjective_text := GetAdjectiveText(*phrase.ObjectCompliment.Adjective)
			text_slice = append(text_slice, adjective_text)
		}
		if phrase.ObjectCompliment.Adjective != nil {
			adjective_text := GetAdjectiveText(*phrase.ObjectCompliment.Adjective)
			text_slice = append(text_slice, adjective_text)
		}
	}

	// Get Text from Prepositional Compliment struct
	if phrase.PrepositionalCompliment != nil {
		if phrase.PrepositionalCompliment.Preposition != nil {
			preposition_text := phrase.PrepositionalCompliment.Preposition.Preposition
			text_slice = append(text_slice, preposition_text)
		}
		if phrase.PrepositionalCompliment.NounPhrase != nil {
			noun_text := GetNounText(*phrase.PrepositionalCompliment.NounPhrase)
			text_slice = append(text_slice, noun_text)
		}
	}

	text := strings.Join(text_slice, " ")

	return text
}

// Gets all text from within a clause
func GetClauseText(clause ClauseType) string {
	text_slice := []string{}

	subject_text := GetNounText(clause.Subject)
	text_slice = append(text_slice, subject_text)
	verb_text := GetVerbText(clause.Verb)
	text_slice = append(text_slice, verb_text)
	if clause.Object != nil {
		object_text := GetNounText(*clause.Object)
		text_slice = append(text_slice, object_text)
	}
	if clause.Compliment != nil {
		compliment_text := GetComplimentText(*clause.Compliment)
		text_slice = append(text_slice, compliment_text)
	}
	if clause.Adjunct != nil {
		adjunct_text := GetAdverbText(*clause.Adjunct)
		text_slice = append(text_slice, adjunct_text)
	}

	text := strings.Join(text_slice, " ")

	return text
}
