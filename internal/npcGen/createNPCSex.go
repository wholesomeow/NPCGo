package npcgen

import (
	"log"
	"math/rand"

	"github.com/wholesomeow/npcGo/internal/npcGen/enums"
)

func CreateSexType(npc_object *NPCBase) error {
	var err error

	log.Print("generating NPC Sex UUID")
	npc_object.SexType.UUID, err = CreateUUIDv4()
	if err != nil {
		return err
	}

	log.Print("selecting NPC Sex")
	sex_select := rand.Intn(3) + 1
	npc_object.SexType.Enum = enums.SexType(sex_select)
	npc_object.SexType.Name = npc_object.SexType.Enum.SexStateToString()

	return nil
}
