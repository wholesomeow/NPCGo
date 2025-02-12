package npc

import (
	"fmt"
	npc "go/npcGen/npc/enums"
)

// TODO(wholesomeow): Implement UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services
type NPCBase struct {
	// v0.1
	UUID int
	Name string
	// OCEAN
	// Enneagram
	// MICE
	// CS
	// REI

	// v0.2
	// Race - Will need to be its own struct
	NPCEnums struct { // Collector for all NPC Enums
		npc.NPCType
		npc.BodyType
		npc.SexType
		npc.GenderType
		npc.OrientationType
	}
	Pronouns      string
	NPCAppearance struct {
		Height_Ft  int
		Height_In  int
		Weight_Lbs int
		Height_Cm  float64
		Weight_Kg  float64
	}

	// v1.0
	// Social Role
	// Communication Matrix
	// Social Circle

	// v1.1
	// Physical Description

	// v1.2
	// Rumors Known
	// Jobs Known
}

func DisplayName(npc NPCBase) {
	fmt.Printf("NPC Name: %s\n", npc.Name)
}
