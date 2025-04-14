package texttypes

import (
	"strings"
)

// Adjectives give us more information about people, animals or things represented by nouns and pronouns.
type AdjectiveType struct {
	Adjective string
	Positive  bool
	Category  string
}

type AdjectivePhrase struct {
	Modifier *AdverbPhrase
	Opinion  *AdjectiveType
	Size     *AdjectiveType
	Quality  *AdjectiveType
	Age      *AdjectiveType
	Shape    *AdjectiveType
	Colour   *AdjectiveType
	Origin   *AdjectiveType
	Material *AdjectiveType
	Type     *AdjectiveType
	Purpose  *AdjectiveType
}

// Builds an Adjective Phrase struct from passed keywords.
// Takes in keywords, ID's the lexical category/POS of keywords
// and assigns them to a phraseType
func (phrase AdjectivePhrase) BuildAdjPhrase(adj AdjectiveType) AdjectivePhrase {
	category := strings.ToLower(adj.Category)
	switch category {
	case "opinion":
		phrase.Opinion = &adj
	case "size":
		phrase.Size = &adj
	case "quality":
		phrase.Quality = &adj
	case "age":
		phrase.Age = &adj
	case "shape":
		phrase.Shape = &adj
	case "colour":
		phrase.Colour = &adj
	case "color":
		phrase.Colour = &adj
	case "origin":
		phrase.Origin = &adj
	case "material":
		phrase.Material = &adj
	case "type":
		phrase.Type = &adj
	case "purpose":
		phrase.Purpose = &adj
	}
	return phrase
}

// Gets all text from within an adjective phrase
func GetAdjectiveText(phrase AdjectivePhrase) string {
	text_slice := []string{}

	if phrase.Modifier != nil {
		modifier_text := GetAdverbText(*phrase.Modifier)
		text_slice = append(text_slice, modifier_text)
	}
	if phrase.Opinion != nil {
		opinon_text := phrase.Opinion.Adjective
		text_slice = append(text_slice, opinon_text)
	}
	if phrase.Size != nil {
		size_text := phrase.Size.Adjective
		text_slice = append(text_slice, size_text)
	}
	if phrase.Quality != nil {
		quality_text := phrase.Quality.Adjective
		text_slice = append(text_slice, quality_text)
	}
	if phrase.Age != nil {
		age_text := phrase.Age.Adjective
		text_slice = append(text_slice, age_text)
	}
	if phrase.Shape != nil {
		shape_text := phrase.Shape.Adjective
		text_slice = append(text_slice, shape_text)
	}
	if phrase.Colour != nil {
		colour_text := phrase.Colour.Adjective
		text_slice = append(text_slice, colour_text)
	}
	if phrase.Origin != nil {
		origin_text := phrase.Origin.Adjective
		text_slice = append(text_slice, origin_text)
	}
	if phrase.Material != nil {
		material_text := phrase.Material.Adjective
		text_slice = append(text_slice, material_text)
	}
	if phrase.Type != nil {
		type_text := phrase.Type.Adjective
		text_slice = append(text_slice, type_text)
	}
	if phrase.Purpose != nil {
		purpose_text := phrase.Purpose.Adjective
		text_slice = append(text_slice, purpose_text)
	}
	text := strings.Join(text_slice, " ")

	return text
}
