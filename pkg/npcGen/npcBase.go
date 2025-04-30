package npcgen

import (
	"encoding/json"

	"github.com/wholesomeow/npcGo/pkg/npcGen/enums"
)

// A NPCBase is the base struct for the NPC Object created by the application
// and contains all of the data to display the NPC descriptions requested
// as well as all the data needed to create some variations
type NPCBase struct {
	UUID     string
	Name     string
	Pronouns []string
	OCEAN    struct {
		UUID        string
		Aspect      []string
		Degree      []float64
		Traits      [][]string
		Text        string
		Description []string
		Use         string
	}
	MICE struct {
		UUID        string
		Aspect      string
		Degree      float64
		Traits      []string
		Text        string
		Description string
		Use         string
	}
	CS struct {
		UUID        string
		Aspect      string
		Traits      []string
		Text        string
		Coords      [2]int
		Description string
		Use         string
	}
	REI struct {
		UUID        string
		Aspect      []string
		Degree      [2]float64
		Traits      []string
		Text        string
		Description []string
		Use         string
	}
	Enneagram struct {
		UUID                  string
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
		UUID        string
		Name        string
		Description string
		Enum        enums.NPCType
	}
	BodyType struct {
		UUID string
		Name string
		Enum enums.BodyType
	}
	SexType struct {
		UUID string
		Name string
		Enum enums.SexType
	}
	GenderType struct {
		UUID        string
		Name        string
		Description string
		Enum        enums.GenderType
	}
	SexualOrientationType struct {
		UUID        string
		Name        string
		Description string
		Enum        enums.OrientationType
	}

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

// Returns all the data within the NPC Object as a JSON object
func (npc_object *NPCBase) DataToJSON() string {
	// result, _ := json.MarshalIndent(npc_object, "", "  ")
	result, _ := json.Marshal(npc_object)

	return string(result)
}

func (npc_object *NPCBase) NameToJSON() string {
	result, _ := json.Marshal(npc_object.Name)

	return string(result)
}

func (npc_object *NPCBase) OCEANToJSON() string {
	// result, _ := json.MarshalIndent(npc_object, "", "  ")
	result, _ := json.Marshal(npc_object.OCEAN)

	return string(result)
}

func (npc_object *NPCBase) MICEToJSON() string {
	// result, _ := json.MarshalIndent(npc_object, "", "  ")
	result, _ := json.Marshal(npc_object.MICE)

	return string(result)
}

func (npc_object *NPCBase) CSToJSON() string {
	// result, _ := json.MarshalIndent(npc_object, "", "  ")
	result, _ := json.Marshal(npc_object.CS)

	return string(result)
}

func (npc_object *NPCBase) REIToJSON() string {
	// result, _ := json.MarshalIndent(npc_object, "", "  ")
	result, _ := json.Marshal(npc_object.REI)

	return string(result)
}

func (npc_object *NPCBase) EnneagramToJSON() string {
	// result, _ := json.MarshalIndent(npc_object, "", "  ")
	result, _ := json.Marshal(npc_object.Enneagram)

	return string(result)
}
