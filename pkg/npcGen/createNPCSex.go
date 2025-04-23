package npcgen

import (
	"log"
	"math/rand"

	"github.com/wholesomeow/npcGo/pkg/npcGen/enums"
)

func (npc_object *NPCBase) CreateSexType() {
	log.Print("selecting NPC Sex")
	sex_select := rand.Intn(3) + 1
	npc_object.SexType.Enum = enums.SexType(sex_select)
	npc_object.SexType.Name = npc_object.SexType.Enum.SexStateToString()
}
