package npcgen

import (
	"log"
	"math/rand"

	"github.com/wholesomeow/npcGo/pkg/npcGen/enums"
)

func (npc_object *NPCBase) CreateOrientationType() error {
	var err error

	log.Print("generating NPC Sexual Orientation UUID")
	npc_object.NPCType.UUID, err = CreateUUIDv4()
	if err != nil {
		return err
	}

	log.Print("selecting NPC Sexual Orientation")
	orientation_select := rand.Intn(len(enums.OriStateName)) + 1
	npc_object.SexualOrientationType.Enum = enums.OrientationType(orientation_select)
	npc_object.SexualOrientationType.Name = npc_object.SexualOrientationType.Enum.OriStateToString()
	npc_object.SexualOrientationType.Description = npc_object.SexualOrientationType.Enum.GetOriDescription()

	return nil
}
