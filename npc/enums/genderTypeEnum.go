package npc

import (
	"fmt"
	"strings"
)

// Enum for the Gender type of an NPC. Returns a unique int as a descriptor of that state.
type GenderType int

var Pronouns = map[int][]string{
	1: {"he", "him"},
	2: {"she", "her"},
	3: {"they", "them"},
}

const (
	Masc_Pronouns    int = 1
	Femme_Pronouns   int = 2
	Neutral_Pronouns int = 3
)

// TODO(wholesomeow): Update these values to contribute to UUID
// const value of enum states
const (
	AGENDER          GenderType = 1
	BIGENDER         GenderType = 2
	CISGENDER        GenderType = 3
	GENDERFLUID      GenderType = 4
	GENDERVARIANT    GenderType = 5
	NONBINARY        GenderType = 6
	TRANSGENDERMAN   GenderType = 7
	TRANSGENDERWOMAN GenderType = 8
)

// map of enum states
var GenStateName = map[GenderType]string{
	AGENDER:          "AGENDER",
	BIGENDER:         "BIGENDER",
	CISGENDER:        "CISGENDER",
	GENDERFLUID:      "GENDERFLUID",
	GENDERVARIANT:    "GENDERVARIANT",
	NONBINARY:        "NONBINARY",
	TRANSGENDERMAN:   "TRANSGENDERMAN",
	TRANSGENDERWOMAN: "TRANSGENDERWOMAN",
}

// string func takes state and returns descriptor
func GenStateToString(gen_state GenderType) string {
	return GenStateName[gen_state]
}

// Checks if the desired state of the enum matches the current state of the enum.
// Returns true if the states do not match, false if they do.
func checkGenTransitionState(gen_state GenderType, desired_state string) bool {
	string_state := GenStateToString(gen_state)
	return string_state != desired_state
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func GenTransition(gen_state GenderType, desired_state string) GenderType {
	desired_state = strings.ToUpper(desired_state)
	switch desired_state {
	case "AGENDER":
		if checkGenTransitionState(gen_state, desired_state) {
			return AGENDER
		}
		return gen_state
	case "BIGENDER":
		if checkGenTransitionState(gen_state, desired_state) {
			return BIGENDER
		}
		return gen_state
	case "CISGENDER":
		if checkGenTransitionState(gen_state, desired_state) {
			return CISGENDER
		}
		return gen_state
	case "GENDERFLUID":
		if checkGenTransitionState(gen_state, desired_state) {
			return GENDERFLUID
		}
		return gen_state
	case "GENDERVARIANT":
		if checkGenTransitionState(gen_state, desired_state) {
			return GENDERVARIANT
		}
		return gen_state
	case "NONBINARY":
		if checkGenTransitionState(gen_state, desired_state) {
			return NONBINARY
		}
		return gen_state
	case "TRANSGENDERMAN":
		if checkGenTransitionState(gen_state, desired_state) {
			return TRANSGENDERMAN
		}
		return gen_state
	case "TRANSGENDERWOMAN":
		if checkGenTransitionState(gen_state, desired_state) {
			return TRANSGENDERWOMAN
		}
		return gen_state
	default:
		panic(fmt.Errorf("unknown state: %s", desired_state))
	}
}
