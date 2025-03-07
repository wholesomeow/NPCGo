package generators

import texttypes "go/npcGen/text_gen/text_types"

type TextData struct {
	Name                  texttypes.NounType
	SubjectivePronoun     texttypes.NounType
	ObjectivePronoun      texttypes.NounType
	PossesstivePronoun    texttypes.NounType
	PositiveAuxiliaryVerb []texttypes.VerbType
	NegativeAuxiliaryVerb []texttypes.VerbType
	Traits                []texttypes.AdjectiveType
	Attributes            []texttypes.AdverbType
	Keywords              []texttypes.AdjectiveType
}
