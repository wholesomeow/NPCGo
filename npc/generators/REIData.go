package generators

import (
	"log"
	"math"
)

// TODO(wholesomeow): Figure out how to use the questionare from this link https://www.psytoolkit.org/survey-library/thinking-style-rei.html
// to *actually* determine Rationality vs Experiential attributes

// TODO(wholesomeow): Figure out what I'm doing with this
func CreateREIAspect(cs_data [2]int) string {
	log.Print("generating Rational-Experiential Inventory Aspect for NPC")
	var aspect string
	aspect_slice := []string{"Rational Ability", "Rational Engagement"}

	// Determine X first
	if cs_data[0] <= 0 {
		aspect_slice[0] = "Experiential Ability"
	}

	// Then Y
	if cs_data[1] <= 0 {
		aspect_slice[1] = "Experiential Engagement"
	}

	return aspect
}

func CreateREIDegree(cs_data [2]int) [2]float64 {
	log.Print("generating Rational-Experiential Inventory Degree for NPC")
	degree := [2]float64{}
	degree[0] = math.Abs(float64(cs_data[0]))
	degree[1] = math.Abs(float64(cs_data[1]))

	return degree
}

func CreateREITraits() []string {
	log.Print("generating Rational-Experiential Inventory Traits for NPC")
	traits := []string{}
	return traits
}

func CreateREIDesc(rei_data [][]string, cs_data [2]int) []string {
	log.Print("populating Rational-Experiential Inventory Description for NPC")
	description := []string{rei_data[0][2], rei_data[1][2]}

	// Determine X first
	if cs_data[0] <= 0 {
		description[0] = rei_data[2][2]
	}

	// Then Y
	if cs_data[1] <= 0 {
		description[1] = rei_data[3][2]
	}

	return description
}

func CreateREIUse() string {
	log.Print("populating Rational-Experiential Inventory Use for NPC")
	return "used to quantify if a person engages in fast intuitive thinking or slow logical thinking"
}
