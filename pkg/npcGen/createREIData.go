package npcgen

import (
	"log"
	"math"
)

// TODO(wholesomeow): Figure out how to use the questionare from this link https://www.psytoolkit.org/survey-library/thinking-style-rei.html
// to *actually* determine Rationality vs Experiential attributes

// TODO(wholesomeow): Figure out what I'm doing with this
func (npc_object *NPCBase) CreateREIData(rei_data [][]string) {
	cs_data := npc_object.CS.Coords

	log.Print("generating Rational-Experiential Inventory Aspect for NPC")
	aspect_slice := []string{"Rational Ability", "Rational Engagement"}

	// Determine X first
	if cs_data[0] <= 0 {
		aspect_slice[0] = "Experiential Ability"
	}

	// Then Y
	if cs_data[1] <= 0 {
		aspect_slice[1] = "Experiential Engagement"
	}
	npc_object.REI.Aspect = aspect_slice

	log.Print("generating Rational-Experiential Inventory Degree for NPC")
	npc_object.REI.Degree[0] = math.Abs(float64(cs_data[0]))
	npc_object.REI.Degree[1] = math.Abs(float64(cs_data[1]))

	// TODO(wholesomeow): Create the logic for this
	log.Print("generating Rational-Experiential Inventory Traits for NPC")
	traits := []string{}
	npc_object.REI.Traits = traits

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

	npc_object.REI.Description = description

	log.Print("populating Rational-Experiential Inventory Use for NPC")
	npc_object.REI.Use = "used to quantify if a person engages in fast intuitive thinking or slow logical thinking"
}
