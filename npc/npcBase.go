package npc

import (
	"fmt"
)

// TODO(wholesomeow): Implement UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services
type NPCBase struct {
	UUID int
	Name string
}

func DisplayName(npc NPCBase) {
	fmt.Printf("NPC Name: %s\n", npc.Name)
}
