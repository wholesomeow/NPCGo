package enums

import (
	"log"
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

// String func takes state and returns descriptor
func (bod_state BodyType) BodStateToString() string {
	return bodStateName[bod_state]
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func BodTransition(bod_state BodyType, desired_state string) BodyType {
	desired_state = strings.ToUpper(desired_state)
	if bod_state.BodStateToString() != desired_state {
		switch desired_state {
		case "SINEWY":
			return SINEWY
		case "LEAN":
			return LEAN
		case "BUFF":
			return BUFF
		case "BUILT":
			return BUILT
		case "THIN":
			return THIN
		case "AVERAGE":
			return AVERAGE
		case "BIGGER":
			return BIGGER
		case "LARGE":
			return FAT
		case "REEDY":
			return FAT
		case "SOFT":
			return FAT
		case "PLUMP":
			return FAT
		case "FAT":
			return FAT
		default:
			log.Printf("unknown state: %s", desired_state)
		}
	}
	return bod_state
}
