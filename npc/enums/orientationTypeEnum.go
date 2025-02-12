package npc

import (
	"fmt"
	"strings"
)

// Enum for the Sexual Orientation type of an NPC. Returns a unique int as a descriptor of that state.
type OrientationType int

// TODO(wholesomeow): Update these values to contribute to UUID
// const value of enum states
const (
	ASEXUAL    OrientationType = 1
	AROMANTIC  OrientationType = 2
	STRAIGHT   OrientationType = 3
	DEMISEXUAL OrientationType = 4
	GAY        OrientationType = 5
	BISEXUAL   OrientationType = 6
	PANSEXUAL  OrientationType = 7
)

// map of enum states with descriptor
var oriStateDesc = map[OrientationType]string{
	ASEXUAL:    "Doesn't often experience sexual attraction.",
	AROMANTIC:  "Doesn't often experience romantic attraction.",
	STRAIGHT:   "Attracted to the sex/gender opposite their own on the spectrum.",
	DEMISEXUAL: "Doesn't experience sexual attraction to someone unless they have a deep, emotional connection with them.",
	GAY:        "Attracted to the sex/gender on the same side of the spectrum.",
	BISEXUAL:   "Attracted to more than one gender or gender identity.",
	PANSEXUAL:  "Attracted to the person rather than their sex, gender, or gender identity.",
}

// map of enum states with descriptor
var OriStateName = map[OrientationType]string{
	ASEXUAL:    "ASEXUAL",
	AROMANTIC:  "AROMANTIC",
	STRAIGHT:   "STRAIGHT",
	DEMISEXUAL: "DEMISEXUAL",
	GAY:        "GAY",
	BISEXUAL:   "BISEXUAL",
	PANSEXUAL:  "PANSEXUAL",
}

// string func takes state and returns string description
func GetOriDescription(ori_state OrientationType) string {
	return oriStateDesc[ori_state]
}

// string func takes state and returns string name
func OriStateToString(ori_state OrientationType) string {
	return OriStateName[ori_state]
}

// Checks if the desired state of the enum matches the current state of the enum.
// Returns true if the states do not match, false if they do.
func checkOriTransitionState(ori_state OrientationType, desired_state string) bool {
	string_state := OriStateToString(ori_state)
	return string_state != desired_state
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func OriTransition(ori_state OrientationType, desired_state string) OrientationType {
	desired_state = strings.ToUpper(desired_state)
	switch desired_state {
	case "ASEXUAL":
		if checkOriTransitionState(ori_state, desired_state) {
			return ASEXUAL
		}
		return ori_state
	case "AROMANTIC":
		if checkOriTransitionState(ori_state, desired_state) {
			return AROMANTIC
		}
		return ori_state
	case "STRAIGHT":
		if checkOriTransitionState(ori_state, desired_state) {
			return STRAIGHT
		}
		return ori_state
	case "DEMISEXUAL":
		if checkOriTransitionState(ori_state, desired_state) {
			return DEMISEXUAL
		}
		return ori_state
	case "GAY":
		if checkOriTransitionState(ori_state, desired_state) {
			return GAY
		}
		return ori_state
	case "BISEXUAL":
		if checkOriTransitionState(ori_state, desired_state) {
			return BISEXUAL
		}
		return ori_state
	case "PANSEXUAL":
		if checkOriTransitionState(ori_state, desired_state) {
			return PANSEXUAL
		}
		return ori_state
	default:
		panic(fmt.Errorf("unknown state: %s", desired_state))
	}
}
