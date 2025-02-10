package npc

import (
	"fmt"
)

type NPCBase struct {
	UUID           int
	Name           string
	Social_Network [3]string //Putting 3 in for now, will need to research for proper implementation
	CS_Dimension   string
	REI_Data       string
}

func (npc NPCBase) DisplayName() {
	fmt.Printf("NPC Name: %s\n", npc.Name)
}
