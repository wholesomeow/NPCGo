package npcgen

import (
	"log"
	"math/rand"

	"github.com/wholesomeow/npcGo/pkg/npcGen/enums"
)

func (npc_object *NPCBase) CreateGenderType() {
	log.Print("selecting NPC Gender")
	gender_select := rand.Intn(len(enums.GenStateName)) + 1
	npc_object.GenderType.Enum = enums.GenderType(gender_select)
	npc_object.GenderType.Name = npc_object.GenderType.Enum.GenStateToString()
	npc_object.GenderType.Description = npc_object.GenderType.Enum.GetGenderDescription()
}

// TODO(wholesomeow): Rework this to allow mixing pronouns
// TODO(wholesomeow): Rework this to be more clear with case to pronoun mapping
func (npc_object *NPCBase) CreatePronouns() {
	log.Print("selecting NPC Pronouns")
	// TODO(wholesomeow): Rework better random selection
	r_val := rand.Intn(len(enums.Pronouns)) + 1
	switch npc_object.GenderType.Enum {
	case 1:
		npc_object.Pronouns = enums.Pronouns[enums.Neutral_Pronouns]
	case 2:
		npc_object.Pronouns = enums.Pronouns[r_val]
	case 3: // TODO(wholesomeow): Figure out how to have sex influence pronoun selection for intersex cisgendered people
		npc_object.Pronouns = enums.Pronouns[r_val]
	case 4: // TODO(wholesomeow): Figure out how gender fluid people prefer to use pronouns
		npc_object.Pronouns = enums.Pronouns[enums.Neutral_Pronouns]
	case 5: // TODO(wholesomeow): Figure out how gender varient people prefer to use pronouns
		npc_object.Pronouns = enums.Pronouns[r_val]
	case 6:
		npc_object.Pronouns = enums.Pronouns[enums.Neutral_Pronouns]
	case 7:
		npc_object.Pronouns = enums.Pronouns[enums.Masc_Pronouns]
	case 8:
		npc_object.Pronouns = enums.Pronouns[enums.Femme_Pronouns]
	}
}
