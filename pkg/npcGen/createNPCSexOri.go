package npcgen

import (
	"go/npcGen/pkg/npcGen/enums"
	"log"
	"math/rand"
)

func (npc_object *NPCBase) CreateOrientationType() {
	log.Print("selecting NPC Sexual Orientation")
	orientation_select := rand.Intn(len(enums.OriStateName)) + 1
	npc_object.SexualOrientationType.Enum = enums.OrientationType(orientation_select)
	npc_object.SexualOrientationType.Name = npc_object.SexualOrientationType.Enum.OriStateToString()
	npc_object.SexualOrientationType.Description = npc_object.SexualOrientationType.Enum.GetOriDescription()
}
