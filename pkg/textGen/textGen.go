package textgen

import (
	texttypes "go/npcGen/pkg/textGen/textTypes"
	"strings"
)

// Takes sentence fragments and conjugates them together using Coordinating
// conjunction words like "and" and "but"
func CoordConj() error {
	return nil
}

// Takes sentence fragments and conjugates them together using Subordinating
// conjunction words like "because" and "if"
func SubordConj() error {
	return nil
}

// Takes sentence fragments and conjugates them together using Temporal
// conjunction words like "while" and "when"
func TemporalConj() error {
	return nil
}

// Builds sentences with one main clause.
// Simple sentences are in Subject, Verb, Object order.
// Object can also be complements or adjunts.
func SimpleSentenceBuilder(data texttypes.TextData) string {
	// Build required Noun/Verb Phrases
	name_phrase := texttypes.NounPhrase{}
	subpn_phrase := texttypes.NounPhrase{}
	posav_phrase := texttypes.VerbPhrase{}
	negav_phrase := texttypes.VerbPhrase{}
	posav_phrase_name := texttypes.VerbPhrase{}
	negav_phrase_name := texttypes.VerbPhrase{}

	name_phrase = texttypes.NounPhrase.BuildNounPhrase(name_phrase, data.Name)
	subpn_phrase = texttypes.NounPhrase.BuildNounPhrase(subpn_phrase, data.SubjectivePronoun)
	posav_phrase_name = texttypes.VerbPhrase.BuildVerbPhrase(posav_phrase_name, data.PositiveAuxiliaryVerb[0])
	negav_phrase_name = texttypes.VerbPhrase.BuildVerbPhrase(negav_phrase_name, data.NegativeAuxiliaryVerb[0])

	if data.SubjectivePronoun.Noun == "they" {
		posav_phrase = texttypes.VerbPhrase.BuildVerbPhrase(posav_phrase, data.PositiveAuxiliaryVerb[1])
		negav_phrase = texttypes.VerbPhrase.BuildVerbPhrase(negav_phrase, data.NegativeAuxiliaryVerb[1])
	} else {
		posav_phrase = texttypes.VerbPhrase.BuildVerbPhrase(posav_phrase, data.PositiveAuxiliaryVerb[0])
		negav_phrase = texttypes.VerbPhrase.BuildVerbPhrase(negav_phrase, data.NegativeAuxiliaryVerb[0])
	}

	clause_slice := []texttypes.ClauseType{}
	attr_count := 0
	for idx, word := range data.Keywords {
		simple_clause := texttypes.ClauseType{}

		// Determine if name or pronoun is used for Subject
		simple_clause.Subject = subpn_phrase
		mod := idx % 3
		if mod == 0 {
			simple_clause.Subject = name_phrase
			attr_count += 1

			// Determine if positive or negative adverbial phrase is used for Verb
			simple_clause.Verb = negav_phrase_name
			if word.Positive {
				simple_clause.Verb = posav_phrase_name
			}
		} else {
			// Determine if positive or negative adverbial phrase is used for Verb
			simple_clause.Verb = negav_phrase
			if word.Positive {
				simple_clause.Verb = posav_phrase
			}
		}

		// Keyword Building
		keyword := texttypes.AdjectivePhrase{}
		attribute := texttypes.AdverbPhrase{}
		attribute.Degree = &data.Attributes[attr_count-1]
		keyword.Modifier = &attribute

		// If not first sentence, keyword is keyword
		if mod != 0 {
			keyword = texttypes.AdjectivePhrase.BuildAdjPhrase(keyword, word)

		} else {
			// Else keyword is trait
			keyword = texttypes.AdjectivePhrase.BuildAdjPhrase(keyword, data.Traits[attr_count-1])
		}

		// Build Subject Compliment
		subject_comp := texttypes.SubjectComplimentType{}
		subject_comp.Adjective = &keyword
		compliment := texttypes.ComplimentType{}
		compliment.SubjectCompliment = &subject_comp
		simple_clause.Compliment = &compliment

		clause_slice = append(clause_slice, simple_clause)
	}

	text_slice := []string{}
	for _, clause := range clause_slice {
		sentence := texttypes.GetClauseText(clause)
		text_slice = append(text_slice, sentence)
	}

	text := strings.Join(text_slice, "\n")

	return text
}

// Builds compound sentences from two or more main clauses, joined by a coordinating conjuction

// Builds complex sentences from one main clause and one or more subordinate clauses, introduced by a subordinating conjunction
