package generators

import (
	"log"
	"math/rand"
)

func CreateMICE(mice_data [][]string) (string, string, string) {
	log.Print("setting MICE values for NPC")
	r_val := rand.Intn(len(mice_data))
	selection := mice_data[r_val]

	aspect := selection[1]
	description := selection[3]
	log.Print("selecting specifc MICE description at index: 3")
	use := "used to list the primary reasons why someone would become a spy, insider threat, or collaborate with a hostile organization"

	return aspect, description, use
}
