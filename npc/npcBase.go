package npc

import (
	"encoding/json"
	npc "go/npcGen/npc/enums"
)

// TODO(wholesomeow): Implement UUID
type NPCBase struct {
	UUID  int
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
		Aspect      string
		Description string
		Use         string
	}
	Enneagram struct {
		ID                    int
		Archetype             string
		Center                string
		DominantEmotion       string
		Keywords              []string
		Description           string
		Fear                  string
		Desire                string
		Wings                 []int
		LODLevel              int
		CurrentLOD            string
		LevelOfDevelopment    []string
		KeyMotivations        string
		Overview              string
		Addictions            string
		GrowthRecommendations []string
	}

	// v1.1
	// Race - Will need to be its own struct
	NPCEnums struct { // Collector for all NPC Enums
		npc.NPCType
		npc.BodyType
		npc.SexType
		npc.GenderType
		npc.OrientationType
	}
	NPCType struct {
		Name        string
		Description string
	}
	BodyType struct {
		Name string
	}
	Sex struct {
		Name string
	}
	Gender struct {
		Name        string
		Description string
	}
	SexualOrientation struct {
		Name        string
		Description string
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

func GetAllNPCData(npc NPCBase) NPCBase {
	return npc
}

func DataToJSON(npc NPCBase) string {
	result, _ := json.MarshalIndent(npc, "", "  ")

	return string(result)
}

func GetName(npc NPCBase) string {
	return npc.Name
}
