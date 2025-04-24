package npcgen

import "log"

func (npc_object *NPCBase) CreateNPCType() error {
	var err error

	log.Print("generating NPC type UUID")
	npc_object.NPCType.UUID, err = CreateUUIDv4()
	if err != nil {
		return err
	}

	npc_object.NPCType.Enum = 0 // Set to DEFAULT on init
	npc_object.NPCType.Name = npc_object.NPCType.Enum.NPCStateToString()
	npc_object.NPCType.Description = npc_object.NPCType.Enum.GetNPCStateDescription()

	return nil
}
