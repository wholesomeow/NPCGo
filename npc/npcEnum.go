package npc

import (
	"fmt"
	"strings"
)

// Enum for game role the NPC takes. Returns a unique int as a descriptor of that state.
type NPCType int
type SexType int
type OrientationType int
type GenderType int
type BodyType int

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
	DEFAULT:   "idle",
	MAIN:      "connected",
	SIDE:      "error",
	IMPORTANT: "retrying",
	RETURNING: "asdf",
	ONEOFF:    "asdf",
	COMPANION: "asdf",
	FRIEND:    "asdf",
	NEUTRAL:   "asdf",
	ENEMY:     "asdf",
	DEAD:      "asdf",
}

// string func takes state and returns descriptor
func StateToString(npc_state NPCType) string {
	return npcStateName[npc_state]
}

// Checks if the desired state of the enum matches the current state of the enum.
// Returns true if the states do not match, false if they do.
func checkTransitionState(npc_state NPCType, desired_state string) bool {
	string_state := StateToString(npc_state)
	return string_state != desired_state
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func transition(npc_state NPCType, desired_state string) NPCType {
	desired_state = strings.ToUpper(desired_state)
	switch desired_state {
	case "DEFAULT":
		if checkTransitionState(npc_state, desired_state) {
			return DEFAULT
		}
		return npc_state
	case "MAIN":
		if checkTransitionState(npc_state, desired_state) {
			return MAIN
		}
		return npc_state
	case "SIDE":
		if checkTransitionState(npc_state, desired_state) {
			return SIDE
		}
		return npc_state
	case "IMPORTANT":
		if checkTransitionState(npc_state, desired_state) {
			return IMPORTANT
		}
		return npc_state
	case "RETURNING":
		if checkTransitionState(npc_state, desired_state) {
			return RETURNING
		}
		return npc_state
	case "ONEOFF":
		if checkTransitionState(npc_state, desired_state) {
			return ONEOFF
		}
		return npc_state
	case "COMPANION":
		if checkTransitionState(npc_state, desired_state) {
			return COMPANION
		}
		return npc_state
	case "FRIEND":
		if checkTransitionState(npc_state, desired_state) {
			return FRIEND
		}
		return npc_state
	case "NEUTRAL":
		if checkTransitionState(npc_state, desired_state) {
			return NEUTRAL
		}
		return npc_state
	case "ENEMY":
		if checkTransitionState(npc_state, desired_state) {
			return ENEMY
		}
		return npc_state
	case "DEAD":
		if checkTransitionState(npc_state, desired_state) {
			return DEAD
		}
		return npc_state
	default:
		panic(fmt.Errorf("unknown state: %s", desired_state))
	}
}
