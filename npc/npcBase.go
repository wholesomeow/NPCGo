package npc

import (
	"encoding/json"
	npc "go/npcGen/npc/enums"
)

// TODO(wholesomeow): Implement UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services
type NPCBase struct {
	UUID      int
	Name      string
	Enneagram Enneagram
	OCEAN     struct {
		Aspect      []float64
		Description []string
		Use         string
	}
	MICE struct {
		Aspect      string
		Description string
		Use         string
	}
	CS struct {
		Aspect      string
		Data        [2]int
		Description string
		Use         string
	}
	REI struct {
		Aspect      string
		Description string
		Use         string
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
	Pronouns      []string
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

	// v2.1
	// Physical Description

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
