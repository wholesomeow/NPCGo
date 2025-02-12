package npc

import (
	"fmt"
	"strings"
)

// Enum for the NPC Body Type of an NPC. Returns a unique int as a descriptor of that state.
type BodyType int

// const value of enum states
const (
	SINEWY  BodyType = 5
	LEAN    BodyType = 10
	BUFF    BodyType = 15
	BUILT   BodyType = 20
	THIN    BodyType = 6
	AVERAGE BodyType = 12
	BIGGER  BodyType = 18
	LARGE   BodyType = 24
	REEDY   BodyType = 7
	SOFT    BodyType = 14
	PLUMP   BodyType = 21
	FAT     BodyType = 28
)

// map of enum states
var bodStateName = map[BodyType]string{
	SINEWY:  "SINEWY",
	LEAN:    "LEAN",
	BUFF:    "BUFF",
	BUILT:   "BUILT",
	THIN:    "THIN",
	AVERAGE: "AVERAGE",
	BIGGER:  "BIGGER",
	LARGE:   "LARGE",
	REEDY:   "REEDY",
	SOFT:    "SOFT",
	PLUMP:   "PLUMP",
	FAT:     "FAT",
}

// string func takes state and returns descriptor
func BodStateToString(bod_state BodyType) string {
	return bodStateName[bod_state]
}

// Checks if the desired state of the enum matches the current state of the enum.
// Returns true if the states do not match, false if they do.
func checkBodTransitionState(bod_state BodyType, desired_state string) bool {
	string_state := BodStateToString(bod_state)
	return string_state != desired_state
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func BodTransition(bod_state BodyType, desired_state string) BodyType {
	desired_state = strings.ToUpper(desired_state)
	switch desired_state {
	case "SINEWY":
		if checkBodTransitionState(bod_state, desired_state) {
			return SINEWY
		}
		return bod_state
	case "LEAN":
		if checkBodTransitionState(bod_state, desired_state) {
			return LEAN
		}
		return bod_state
	case "BUFF":
		if checkBodTransitionState(bod_state, desired_state) {
			return BUFF
		}
		return bod_state
	case "BUILT":
		if checkBodTransitionState(bod_state, desired_state) {
			return BUILT
		}
		return bod_state
	case "THIN":
		if checkBodTransitionState(bod_state, desired_state) {
			return THIN
		}
		return bod_state
	case "AVERAGE":
		if checkBodTransitionState(bod_state, desired_state) {
			return AVERAGE
		}
		return bod_state
	case "BIGGER":
		if checkBodTransitionState(bod_state, desired_state) {
			return BIGGER
		}
		return bod_state
	case "LARGE":
		if checkBodTransitionState(bod_state, desired_state) {
			return FAT
		}
		return bod_state
	case "REEDY":
		if checkBodTransitionState(bod_state, desired_state) {
			return FAT
		}
		return bod_state
	case "SOFT":
		if checkBodTransitionState(bod_state, desired_state) {
			return FAT
		}
		return bod_state
	case "PLUMP":
		if checkBodTransitionState(bod_state, desired_state) {
			return FAT
		}
		return bod_state
	case "FAT":
		if checkBodTransitionState(bod_state, desired_state) {
			return FAT
		}
		return bod_state
	default:
		panic(fmt.Errorf("unknown state: %s", desired_state))
	}
}
