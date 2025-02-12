package npc

import (
	"fmt"
	"strings"
)

// Enum for the Sex type of an NPC. Returns a unique int as a descriptor of that state.
type SexType int

// TODO(wholesomeow): I'm sure there's more depth to be added here for intersex people and others I'm not aware of
// const value of enum states
const (
	MALE   SexType = 1
	FEMALE SexType = 2
	OTHER  SexType = 3
)

// map of enum states
var sexStateName = map[SexType]string{
	MALE:   "MALE",
	FEMALE: "FEMALE",
	OTHER:  "OTHER",
}

// string func takes state and returns string name
func SexStateToString(sex_state SexType) string {
	return sexStateName[sex_state]
}

// Checks if the desired state of the enum matches the current state of the enum.
// Returns true if the states do not match, false if they do.
func checkSexTransitionState(sex_state SexType, desired_state string) bool {
	string_state := SexStateToString(sex_state)
	return string_state != desired_state
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func SexTransition(sex_state SexType, desired_state string) SexType {
	desired_state = strings.ToUpper(desired_state)
	switch desired_state {
	case "MALE":
		if checkSexTransitionState(sex_state, desired_state) {
			return MALE
		}
		return sex_state
	case "FEMALE":
		if checkSexTransitionState(sex_state, desired_state) {
			return FEMALE
		}
		return sex_state
	case "OTHER":
		if checkSexTransitionState(sex_state, desired_state) {
			return OTHER
		}
		return sex_state
	default:
		panic(fmt.Errorf("unknown state: %s", desired_state))
	}
}
