package texttypes

type TextData struct {
	Name                  NounType
	SubjectivePronoun     NounType
	ObjectivePronoun      NounType
	PossesstivePronoun    NounType
	PositiveAuxiliaryVerb []VerbType
	NegativeAuxiliaryVerb []VerbType
	Traits                []AdjectiveType
	Attributes            []AdverbType
	Keywords              []AdjectiveType
}
