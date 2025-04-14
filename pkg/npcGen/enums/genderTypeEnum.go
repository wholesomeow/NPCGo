package enums

import (
	"log"
	"strings"
)

// Enum for the Gender type of an NPC. Returns a unique int as a descriptor of that state.
type GenderType int

var Pronouns = map[int][]string{
	1: {"he", "him", "his"},
	2: {"she", "her", "hers"},
	3: {"they", "them", "theirs"},
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
var genStateDesc = map[GenderType]string{
	AGENDER:          "Not having a gender or identifying with a gender.",
	BIGENDER:         "A person who fluctuates between traditionally “male” and “female” gender-based behaviours and identities.",
	CISGENDER:        "A person whose gender identity and biological sex assigned at birth are the same.",
	GENDERFLUID:      "A person who is gender fluid may always feel like a mix of the two traditional genders but may feel more man some days, and more woman other days.",
	GENDERVARIANT:    "Someone who either by nature or by choice does not conform to gender-based expectations of society.",
	NONBINARY:        "A gender identity and an umbrella term for people whose identity falls outside of the binary of male and female.",
	TRANSGENDERMAN:   "A term used to describe someone who is assigned female at birth but identifies and lives as a man.",
	TRANSGENDERWOMAN: "A term used to describe someone who is assigned male at birth but identifies and lives as a woman.",
}

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

// string func takes state and returns string descriptor
func (gen_state GenderType) GetGenderDescription() string {
	return genStateDesc[gen_state]
}

// string func takes state and returns descriptor
func (gen_state GenderType) GenStateToString() string {
	return GenStateName[gen_state]
}

// Transitions current state of the enum to a specific desired state.
// Takes in a string and the enum and returns the updated enum.
// Returns enum in it's original state if desired state cannot be transitioned to.
func GenTransition(gen_state GenderType, desired_state string) GenderType {
	desired_state = strings.ToUpper(desired_state)
	if gen_state.GenStateToString() != desired_state {
		switch desired_state {
		case "AGENDER":
			return AGENDER
		case "BIGENDER":
			return BIGENDER
		case "CISGENDER":
			return CISGENDER
		case "GENDERFLUID":
			return GENDERFLUID
		case "GENDERVARIANT":
			return GENDERVARIANT
		case "NONBINARY":
			return NONBINARY
		case "TRANSGENDERMAN":
			return TRANSGENDERMAN
		case "TRANSGENDERWOMAN":
			return TRANSGENDERWOMAN
		default:
			log.Printf("unknown state: %s", desired_state)
		}
	}
	return gen_state
}
