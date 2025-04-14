package enums

import (
	"log"
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
func (sex_state SexType) SexStateToString() string {
	return sexStateName[sex_state]
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func SexTransition(sex_state SexType, desired_state string) SexType {
	desired_state = strings.ToUpper(desired_state)
	if sex_state.SexStateToString() != desired_state {
		switch desired_state {
		case "MALE":
			return MALE
		case "FEMALE":
			return FEMALE
		case "OTHER":
			return OTHER
		default:
			log.Printf("unknown state: %s", desired_state)
		}
	}
	return sex_state
}
