package generators

import (
	"log"
	"math/rand"
)

func CreateCSData(cs_data [][]string) (string, [2]int, string, string) {
	log.Print("generating Cognitive Science data for NPC")
	var cs_coords = [2]int{0, 0}
	var selection = []string{}

	min := -100
	max := 100
	cs_coords[0] = rand.Intn((max - min + 1)) + min
	cs_coords[1] = rand.Intn((max - min + 1)) + min

	if cs_coords[0] <= 0 && cs_coords[1] <= 0 {
		selection = cs_data[0]
	} else if cs_coords[0] <= 0 && cs_coords[1] >= 0 {
		selection = cs_data[1]
	} else if cs_coords[0] >= 0 && cs_coords[1] >= 0 {
		selection = cs_data[2]
	} else if cs_coords[0] >= 0 && cs_coords[1] <= 0 {
		selection = cs_data[3]
	}

	aspect := selection[1]
	description := selection[3]
	use := "used to quantify at which cognitive aspects a person either excels at, struggles with, or a combination of both"

	return aspect, cs_coords, description, use
}
