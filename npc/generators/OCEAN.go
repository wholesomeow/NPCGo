package generators

import (
	texttypes "go/npcGen/text_gen/text_types"
	"log"
	"math"
	"strconv"
	"strings"
)

func remapOCEAN(value float64, minInput float64, maxInput float64, minOutput float64, maxOutput float64) float64 {
	var part_1 float64
	var part_2 float64
	part_1 = (value - minInput) / (maxInput - minInput)
	part_2 = (maxOutput - minOutput) + minOutput
	return part_1 * part_2
}

func CreateOCEANData(ocean_data [][]string, cs_data [2]int) ([]float64, [][]string, []string, string) {
	log.Print("generating OCEAN values for NPC")
	aspect := []float64{}

	for _, val := range ocean_data {
		ocean_cast := []float64{}
		// X Coord cast first
		split := strings.Split(string(val[2]), ",")
		x, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			log.Fatalf("Error converting string to integer: %s", err)
		}
		ocean_cast = append(ocean_cast, float64(x))

		// Y Coord cast second
		y, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			log.Fatalf("Error converting string to integer: %s", err)
		}
		ocean_cast = append(ocean_cast, float64(y))

		// Variable casting
		x1 := ocean_cast[0]
		y1 := ocean_cast[1]
		x2 := float64(cs_data[0])
		y2 := float64(cs_data[1])

		out := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
		remapped_out := remapOCEAN(out, -250, 250, -100, 100)
		aspect = append(aspect, remapped_out)
	}

	traits := [][]string{}
	traits = append(traits, []string{"willing to try new things", "think outside the box", "curious", "creative", "imaginative"})
	traits = append(traits, []string{"organized", "thoughtful", "goal-orientated", "disciplined", "persistent"})
	traits = append(traits, []string{"sociable", "assertive", "energy", "talkative", "outgoing"})
	traits = append(traits, []string{"kind", "altruistic", "trusting", "cooperative", "prosocial"})
	traits = append(traits, []string{"anxious", "guilty", "angry", "sullen", "depressed"})

	description := []string{}
	description = append(description, "A person's willingness to try new things and think outside the box. These people are curious, creative, and imaginative.")
	description = append(description, "A person's level of organization, thoughtfulness, and goal-orientation. These people are more disciplined and persistent.")
	description = append(description, "A person's level of sociability, assertiveness, and energy. These people are more likely to be talkative, outgoing, and have a wide social circle.")
	description = append(description, "A person's level of kindness, altruism, and trust. These people are more cooperative and prosocial.")
	description = append(description, "A person's tendency to experience negative emotions like anxiety, guilt, anger, and depression. These people are more likely to experience these feelings.")

	use := "used to broadly describe and analyze a person's personality by identifying five key dimensions of their behavior"

	return aspect, traits, description, use
}

func CreateOCEANText(npc_name string, pronouns []string, traits [][]string, aspect []float64) TextData {
	log.Print("start of OCEAN Text Generation")

	trait_name := []string{"open", "conscientious", "extraverted", "agreeable", "neurotic"}
	trait_slice := []texttypes.AdjectiveType{}
	adj_slice := []texttypes.AdjectiveType{}
	attr_slice := []texttypes.AdverbType{}
	attribute_values := []string{
		"not at all",
		"slightly",
		"somewhat",
		"moderately",
		"fairly",
		"quite",
		"strongly",
		"very",
		"very much",
		"extremely",
	}

	// Word Type declaration
	name := texttypes.NounType{}
	subjectivePN := texttypes.NounType{}
	objectivePN := texttypes.NounType{}
	possessivePN := texttypes.NounType{}
	posAuxVerb_1 := texttypes.VerbType{}
	posAuxVerb_2 := texttypes.VerbType{}
	negAuxVerb_1 := texttypes.VerbType{}
	negAuxVerb_2 := texttypes.VerbType{}

	// Word Type assignment
	name.Noun = npc_name
	subjectivePN.Noun = pronouns[0]
	objectivePN.Noun = pronouns[1]
	possessivePN.Noun = pronouns[2]

	// Assign Words to text data struct
	oceanTextData := TextData{}
	oceanTextData.Name = name
	oceanTextData.SubjectivePronoun = subjectivePN
	oceanTextData.ObjectivePronoun = objectivePN
	oceanTextData.PossesstivePronoun = possessivePN

	log.Print("setting pronoun auxiliary verbs")
	pnoun_aux_verbs := [][]string{{"is", "isn't"}, {"are", "aren't"}}
	posAuxVerb_1.Verb = pnoun_aux_verbs[0][0]
	negAuxVerb_1.Verb = pnoun_aux_verbs[0][1]
	posAuxVerb_2.Verb = pnoun_aux_verbs[1][0]
	negAuxVerb_2.Verb = pnoun_aux_verbs[1][1]

	oceanTextData.PositiveAuxiliaryVerb = append(oceanTextData.PositiveAuxiliaryVerb, posAuxVerb_1)
	oceanTextData.PositiveAuxiliaryVerb = append(oceanTextData.PositiveAuxiliaryVerb, posAuxVerb_2)
	oceanTextData.NegativeAuxiliaryVerb = append(oceanTextData.NegativeAuxiliaryVerb, negAuxVerb_1)
	oceanTextData.NegativeAuxiliaryVerb = append(oceanTextData.NegativeAuxiliaryVerb, negAuxVerb_2)

	// Cycle through all OCEAN values to create text data
	for i := range 5 {
		log.Printf("generating keyword data for trait: %s", trait_name[i])
		trait := texttypes.AdjectiveType{}

		// Determine attribute string from Aspect value
		// TODO(wholeomeow): Implement fuzzing engine for this to expand attribute values
		for j := 0.0; j < 110.0; j += 10.0 {
			var attribute_count float64
			var k float64 = j
			if j != 0.0 {
				attribute_count = j / 10.0
				k -= 10.0
			} else {
				attribute_count = 0.0
				k = 0.0
			}

			positive := false
			if j >= 50.0 {
				positive = true
			}

			if aspect[i] > k && aspect[i] < j {
				log.Printf("match found for OCEAN aspect: %s", trait_name[i])

				trait.Adjective = trait_name[i]
				trait.Positive = positive
				trait.Category = "Quality"
				trait_slice = append(trait_slice, trait)

				attribute := texttypes.AdverbType{}
				attribute.Adverb = attribute_values[int(attribute_count)-1]
				attr_slice = append(attr_slice, attribute)

				// Build long and short trait descriptors
				// TODO(wholesomeow): Implement OCEAN values to lexicon search for more traits
				// long_traits = traits[i][:2] // Comment these out for now
				short_traits := traits[i][2:]
				for _, trait := range short_traits {
					new_adj := texttypes.AdjectiveType{}
					new_adj.Adjective = trait
					new_adj.Positive = positive
					new_adj.Category = "Quality"
					adj_slice = append(adj_slice, new_adj)
				}
			}
		}
	}
	oceanTextData.Traits = trait_slice
	oceanTextData.Attributes = attr_slice
	oceanTextData.Keywords = adj_slice

	return oceanTextData
}
