package generators

import (
	"go/npcGen/utilities"
	"log"
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

func CreateCSAspect(cs_data [][]string, cs_coords [2]int) string {
	log.Print("generating Cognitive Science Aspect for NPC")
	idx := coordsToSelection(cs_coords)
	selection := cs_data[idx]
	return selection[1]
}

func CreateCSCoords(cs_data [][]string) [2]int {
	log.Print("generating Cognitive Science Coordiantes for NPC")
	var cs_coords = [2]int{0, 0}

	cs_coords[0] = utilities.RandomRange(-100, 100)
	cs_coords[1] = utilities.RandomRange(-100, 100)

	return cs_coords
}

func CreateCSTraits(cs_data [][]string, cs_coords [2]int) []string {
	log.Print("generating Cognitive Science Traits for NPC")
	traits := []string{}
	return traits
}

func CreateCSDesc(cs_data [][]string, cs_coords [2]int) string {
	log.Print("populating Cognitive Science Description for NPC")
	idx := coordsToSelection(cs_coords)
	selection := cs_data[idx]
	return selection[3]
}

func CreateCSUse() string {
	log.Print("populating Cognitive Science Use for NPC")
	return "used to quantify at which cognitive aspects a person either excels at, struggles with, or a combination of both"
}
