package npcgen

import (
	"log"

	utilities "github.com/wholesomeow/npcGo/internal/utilities"
)

func coordsToSelection(cs_coords [2]int) int {
	var selection int
	if cs_coords[0] <= 0 && cs_coords[1] <= 0 {
		selection = 0
	} else if cs_coords[0] <= 0 && cs_coords[1] >= 0 {
		selection = 1
	} else if cs_coords[0] >= 0 && cs_coords[1] >= 0 {
		selection = 2
	} else if cs_coords[0] >= 0 && cs_coords[1] <= 0 {
		selection = 3
	}

	return selection
}

func (npc_object *NPCBase) CreateCSData(cs_data [][]string) {
	log.Print("generating Cognitive Science Aspect for NPC")
	idx := coordsToSelection(npc_object.CS.Coords)
	selection := cs_data[idx]
	npc_object.CS.Aspect = selection[1]

	log.Print("generating Cognitive Science Coordiantes for NPC")
	npc_object.CS.Coords[0] = utilities.RandomRange(-100, 100)
	npc_object.CS.Coords[1] = utilities.RandomRange(-100, 100)

	// TODO(wholesomeow): Create the logic for this
	log.Print("generating Cognitive Science Traits for NPC")
	traits := []string{}
	npc_object.CS.Traits = traits

	log.Print("populating Cognitive Science Description for NPC")
	npc_object.CS.Description = selection[2]

	log.Print("populating Cognitive Science Use for NPC")
	npc_object.CS.Use = "used to quantify at which cognitive aspects a person either excels at, struggles with, or a combination of both"
}
