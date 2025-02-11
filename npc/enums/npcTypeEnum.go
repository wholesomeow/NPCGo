package npc

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
var npcStateName = map[NPCType]string{
	DEFAULT:   "Default descriptor",
	MAIN:      "Main descriptor",
	SIDE:      "Side descriptor",
	IMPORTANT: "Important descriptor",
	RETURNING: "Returning descriptor",
	ONEOFF:    "One-off descriptor",
	COMPANION: "Companion descriptor",
	FRIEND:    "Friend descriptor",
	NEUTRAL:   "Neutral descriptor",
	ENEMY:     "Enemy descriptor",
	DEAD:      "Dead descriptor",
}

// string func takes state and returns descriptor
func NPCStateToString(npc_state NPCType) string {
	return npcStateName[npc_state]
}

// Checks if the desired state of the enum matches the current state of the enum.
// Returns true if the states do not match, false if they do.
func checkNPCTransitionState(state NPCType, desired_state string) bool {
	string_state := NPCStateToString(state)
	return string_state != desired_state
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func npcTransition(state NPCType, desired_state string) NPCType {
	desired_state = strings.ToUpper(desired_state)
	switch desired_state {
	case "DEFAULT":
		if checkNPCTransitionState(state, desired_state) {
			return DEFAULT
		}
		return state
	case "MAIN":
		if checkNPCTransitionState(state, desired_state) {
			return MAIN
		}
		return state
	case "SIDE":
		if checkNPCTransitionState(state, desired_state) {
			return SIDE
		}
		return state
	case "IMPORTANT":
		if checkNPCTransitionState(state, desired_state) {
			return IMPORTANT
		}
		return state
	case "RETURNING":
		if checkNPCTransitionState(state, desired_state) {
			return RETURNING
		}
		return state
	case "ONEOFF":
		if checkNPCTransitionState(state, desired_state) {
			return ONEOFF
		}
		return state
	case "COMPANION":
		if checkNPCTransitionState(state, desired_state) {
			return COMPANION
		}
		return state
	case "FRIEND":
		if checkNPCTransitionState(state, desired_state) {
			return FRIEND
		}
		return state
	case "NEUTRAL":
		if checkNPCTransitionState(state, desired_state) {
			return NEUTRAL
		}
		return state
	case "ENEMY":
		if checkNPCTransitionState(state, desired_state) {
			return ENEMY
		}
		return state
	case "DEAD":
		if checkNPCTransitionState(state, desired_state) {
			return DEAD
		}
		return state
	default:
		panic(fmt.Errorf("unknown state: %s", desired_state))
	}
}
