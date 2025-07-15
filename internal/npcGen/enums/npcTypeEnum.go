package enums

import (
	"fmt"
	"strings"
)

// Enum for game role the NPC takes. Returns a unique int as a descriptor of that state.
type NPCType int

// const values of enum states
const (
	DEFAULT   NPCType = 0
	MAIN      NPCType = 1
	SIDE      NPCType = 2
	IMPORTANT NPCType = 4
	RETURNING NPCType = 8
	ONEOFF    NPCType = 16
	COMPANION NPCType = 32
	FRIEND    NPCType = 64
	NEUTRAL   NPCType = 128
	ENEMY     NPCType = 256
	DEAD      NPCType = 512
)

// map of enum states with descriptor
var npcStateDesc = map[NPCType]string{
	DEFAULT:   "Default NPC State",
	MAIN:      "This NPC is a main character in the plot",
	SIDE:      "This NPC is a side character in the plot",
	IMPORTANT: "This NPC is a important character in the plot",
	RETURNING: "This NPC is a returning character in the plot",
	ONEOFF:    "This NPC is a one-off character in the plot",
	COMPANION: "This NPC is a companion to the party",
	FRIEND:    "This NPC is a friend to the party",
	NEUTRAL:   "This NPC is neutral to the party",
	ENEMY:     "This NPC is an enemy to the party",
	DEAD:      "This NPC is dead",
}

// map of enum states with descriptor
var npcStateName = map[NPCType]string{
	DEFAULT:   "DEFAULT",
	MAIN:      "MAIN",
	SIDE:      "SIDE",
	IMPORTANT: "IMPORTANT",
	RETURNING: "RETURNING",
	ONEOFF:    "ONEOFF",
	COMPANION: "COMPANION",
	FRIEND:    "FRIEND",
	NEUTRAL:   "NEUTRAL",
	ENEMY:     "ENEMY",
	DEAD:      "DEAD",
}

// String func takes state and returns enum descriptor
func (npc_state NPCType) GetNPCStateDescription() string {
	return npcStateDesc[npc_state]
}

// String func takes state and returns enum name
func (npc_state NPCType) NPCStateToString() string {
	return npcStateName[npc_state]
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func NPCTransition(state NPCType, desired_state string) NPCType {
	desired_state = strings.ToUpper(desired_state)
	if state.NPCStateToString() != desired_state {
		switch desired_state {
		case "DEFAULT":
			return DEFAULT
		case "MAIN":
			return MAIN
		case "SIDE":
			return SIDE
		case "IMPORTANT":
			return IMPORTANT
		case "RETURNING":
			return RETURNING
		case "ONEOFF":
			return ONEOFF
		case "COMPANION":
			return COMPANION
		case "FRIEND":
			return FRIEND
		case "NEUTRAL":
			return NEUTRAL
		case "ENEMY":
			return ENEMY
		case "DEAD":
			return DEAD
		default:
			panic(fmt.Errorf("unknown state: %s", desired_state))
		}
	}
	return state
}
