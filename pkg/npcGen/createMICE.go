package npcgen

import (
	"log"
	"math"
	"strconv"
	"strings"

	utilities "github.com/wholesomeow/npcGo/internal/utilities"
)

func (npc_object *NPCBase) CreateMICEData(r_val int, mice_data [][]string) {
	cs_data := npc_object.CS.Coords

	log.Print("setting MICE values for NPC")
	selection := mice_data[r_val]
	npc_object.MICE.Aspect = selection[1]

	log.Print("creating MICE Degree value")
	mice_cast := []float64{}
	// X Coord cast first
	split := strings.Split(string(selection[1]), ",")
	x, err := strconv.Atoi(strings.TrimSpace(split[0]))
	if err != nil {
		log.Fatalf("Error converting string to X coordinate integer: %s", err)
	}
	mice_cast = append(mice_cast, float64(x))

	// Y Coord cast second
	y, err := strconv.Atoi(strings.TrimSpace(split[1]))
	if err != nil {
		log.Fatalf("Error converting string to Y coordinate integer: %s", err)
	}
	mice_cast = append(mice_cast, float64(y))

	// Variable casting
	x1 := mice_cast[0]
	y1 := mice_cast[1]
	x2 := float64(cs_data[0])
	y2 := float64(cs_data[1])

	out := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	npc_object.MICE.Degree = utilities.RemapInt(out, -250, 250, -100, 100)

	// TODO(wholesomeow): Create the logic for this
	// TODO(wholesomeow): Figure out how I'm going to create a traits list to describe
	// how someone could be convinced/manipulated
	log.Print("creating MICE Traits")
	traits := []string{}
	npc_object.MICE.Traits = traits

	log.Print("setting MICE Description")
	log.Print("selecting specifc MICE description at index: 3")
	npc_object.MICE.Description = selection[2]

	log.Print("setting MICE Usage")
	npc_object.MICE.Use = "used to list the primary reasons why someone would become a spy, insider threat, or collaborate with a hostile organization"
}
