package generators

import (
	"go/npcGen/utilities"
	"log"
	"math"
	"strconv"
	"strings"
)

func CreateMICEAspect(r_val int, mice_data [][]string, cs_data [2]int) string {
	log.Print("setting MICE values for NPC")
	selection := mice_data[r_val]
	aspect := selection[1]

	return aspect
}

func CreateMICEDegree(r_val int, mice_data [][]string, cs_data [2]int) float64 {
	selection := mice_data[r_val]

	mice_cast := []float64{}
	// X Coord cast first
	split := strings.Split(string(selection[2]), ",")
	x, err := strconv.Atoi(strings.TrimSpace(split[0]))
	if err != nil {
		log.Fatalf("Error converting string to integer: %s", err)
	}
	mice_cast = append(mice_cast, float64(x))

	// Y Coord cast second
	y, err := strconv.Atoi(strings.TrimSpace(split[1]))
	if err != nil {
		log.Fatalf("Error converting string to integer: %s", err)
	}
	mice_cast = append(mice_cast, float64(y))

	// Variable casting
	x1 := mice_cast[0]
	y1 := mice_cast[1]
	x2 := float64(cs_data[0])
	y2 := float64(cs_data[1])

	out := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	degree := utilities.RemapInt(out, -250, 250, -100, 100)

	return degree
}

func CreateMICETraits(r_val int, mice_data [][]string, cs_data [2]int) []string {
	// TODO(wholesomeow): Figure out how I'm going to create a traits list to describe
	// how someone could be convinced/manipulated
	traits := []string{}
	return traits
}

func CreateMICEDesc(r_val int, mice_data [][]string, cs_data [2]int) string {
	selection := mice_data[r_val]
	description := selection[3]
	log.Print("selecting specifc MICE description at index: 3")
	return description
}

func CreateMICEUse() string {
	return "used to list the primary reasons why someone would become a spy, insider threat, or collaborate with a hostile organization"
}
