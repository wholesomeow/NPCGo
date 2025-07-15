package texttypes

import "strings"

// Adverbs add more information about a verb, an adjective, another adverb, a clause or a whole sentence.
type AdverbType struct {
	Adverb string
}

// Manner: How something happens.
// Place: Where something happens.
// Time: When something happens.
// Duration: How long something happens.
// Frequency: How often something happens.
// Focus: Something specific.
// Degree: To what degree something happens.
// Certainty: How certain something is.
// Evaluation: Speakers opinion of something.
// Perspective: Speakers perspective or reaction.
// Linking: Relationships between clauses and sentences.
type AdverbPhrase struct {
	Manner      *AdverbType
	Place       *AdverbType
	Time        *AdverbType
	Duration    *AdverbType
	Frequency   *AdverbType
	Focus       *AdverbType
	Degree      *AdverbType
	Certainty   *AdverbType
	Evaluation  *AdverbType
	Perspective *AdverbType
	Linking     *AdverbType
}

// Gets all text from within an adverb phrase
func GetAdverbText(phrase AdverbPhrase) string {
	text_slice := []string{}

	if phrase.Manner != nil {
		manner_text := phrase.Manner.Adverb
		text_slice = append(text_slice, manner_text)
	}
	if phrase.Place != nil {
		place_text := phrase.Place.Adverb
		text_slice = append(text_slice, place_text)
	}
	if phrase.Time != nil {
		time_text := phrase.Time.Adverb
		text_slice = append(text_slice, time_text)
	}
	if phrase.Duration != nil {
		duration_text := phrase.Duration.Adverb
		text_slice = append(text_slice, duration_text)
	}
	if phrase.Frequency != nil {
		frequency_text := phrase.Frequency.Adverb
		text_slice = append(text_slice, frequency_text)
	}
	if phrase.Focus != nil {
		focus_text := phrase.Focus.Adverb
		text_slice = append(text_slice, focus_text)
	}
	if phrase.Degree != nil {
		degree_text := phrase.Degree.Adverb
		text_slice = append(text_slice, degree_text)
	}
	if phrase.Certainty != nil {
		certainty_text := phrase.Certainty.Adverb
		text_slice = append(text_slice, certainty_text)
	}
	if phrase.Evaluation != nil {
		evaluation_text := phrase.Evaluation.Adverb
		text_slice = append(text_slice, evaluation_text)
	}
	if phrase.Perspective != nil {
		perspective_text := phrase.Perspective.Adverb
		text_slice = append(text_slice, perspective_text)
	}
	if phrase.Linking != nil {
		linking_text := phrase.Linking.Adverb
		text_slice = append(text_slice, linking_text)
	}
	text := strings.Join(text_slice, " ")

	return text
}
