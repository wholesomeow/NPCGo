package npcgen

import (
	"encoding/json"

	"github.com/wholesomeow/npcGo/pkg/npcGen/enums"
)

// A NPCBase is the base struct for the NPC Object created by the application
// and contains all of the data to display the NPC descriptions requested
// as well as all the data needed to create some variations
type NPCBase struct {
	// TODO(wholesomeow): Implement UUID
	UUID  string
	Name  string
	OCEAN struct {
		Aspect      []string
		Degree      []float64
		Traits      [][]string
		Text        string
		Description []string
		Use         string
	}
	MICE struct {
		Aspect      string
		Degree      float64
		Traits      []string
		Text        string
		Description string
		Use         string
	}
	CS struct {
		Aspect      string
		Traits      []string
		Text        string
		Coords      [2]int
		Description string
		Use         string
	}
	REI struct {
		Aspect      []string
		Degree      [2]float64
		Traits      []string
		Text        string
		Description []string
		Use         string
	}
	Enneagram struct {
		ID                    int
		Archetype             string
		Keywords              []string
		Description           string
		Center                string
		DominantEmotion       string
		Fear                  string
		Desire                string
		Wings                 [2]int
		LODLevel              int
		CurrentLOD            string
		LevelOfDevelopment    [9]string
		KeyMotivations        string
		Overview              string
		Addictions            string
		GrowthRecommendations [5]string
	}

	// v1.1
	// Race - Will need to be its own struct
	NPCType struct {
		Name        string
		Description string
		Enum        enums.NPCType
	}
	BodyType struct {
		Name string
		Enum enums.BodyType
	}
	SexType struct {
		Name string
		Enum enums.SexType
	}
	GenderType struct {
		Name        string
		Description string
		Enum        enums.GenderType
	}
	SexualOrientationType struct {
		Name        string
		Description string
		Enum        enums.OrientationType
	}
	Pronouns []string
	// v2.1
	// Physical Description
	NPCAppearance struct {
		Height_Ft  int
		Height_In  int
		Weight_Lbs int
		Height_Cm  float64
		Weight_Kg  float64
	}

	// v2.0
	// Social Role
	// Communication Matrix
	// Social Circle

	// v2.2
	// Rumors Known
	// Jobs Known
}

func (npc_object *NPCBase) DataToJSON() string {
	result, _ := json.MarshalIndent(npc_object, "", "  ")

	return string(result)
}
