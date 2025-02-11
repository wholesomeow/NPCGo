package npc

import (
	"fmt"
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
	// Type
	// Race
	// Sex
	// Gender
	// Pronouns
	// Sexual Orientation
	// Body Type
	// Height and Weight

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
