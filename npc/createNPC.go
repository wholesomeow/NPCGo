package npc

import (
	"encoding/json"
	"fmt"
	"go/npcGen/configuration"
	npc "go/npcGen/npc/enums"
	"go/npcGen/utilities"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// --------------------------------------------------- CREATE NPC NAME BEGIN ---------------------------------------------------
func CreateName(config *configuration.Config) string {
	var mchain MarkovChain
	var name string
	max_attempts := 6

	buildNGram(&mchain, config, max_attempts)

	log.Print("starting name creation")
	start_proc := time.Now()
	for count := range max_attempts {
		log.Printf("name creation attempt %d", count)
		name = makeName(&mchain)
		if checkQuality(&mchain, name) {
			break
		}
		log.Printf("name %s doesn't meet quality check... moving on to next attempt", name)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("name creation completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	return name
}

// --------------------------------------------------- CREATE NPC BODY BEGIN ---------------------------------------------------
// TODO(wholesomeow): There's probably a better way to implement these values here
func makeBMI(BMI float64) int {
	log.Print("creating NPC BMI")
	if BMI <= 18.5 {
		return 1
	} else if 18.5 < BMI || BMI <= 24.9 {
		return 2
	} else if 25 < BMI || BMI <= 29.29 {
		return 3
	} else {
		return 4
	}
}

// TODO(wholesomeow): There's probably a better way to implement these values here
func MakeSizeImperial() (int, int, int, int) {
	ft := 0
	inch := 0
	lbs := 0

	ft_medium := []int{4, 5, 6, 7}
	ft_small := []int{2, 3}

	lbs_min := 110
	lbs_max := 250

	r_height := rand.Intn(2)
	if r_height == 0 {
		ft = ft_medium[rand.Intn(len(ft_medium))]
		inch = rand.Intn(11)
		lbs = rand.Intn(lbs_max-lbs_min+1) + lbs_min
	} else {
		ft = ft_small[rand.Intn(len(ft_small))]
		inch = rand.Intn(11)
		lbs = rand.Intn(lbs_max-lbs_min+1) + lbs_min
	}

	inches := (ft * 12) + inch

	return ft, inch, lbs, inches
}

func MakeSizeMetric(inches int, lbs int) (float64, float64) {
	return utilities.ImperialToMetric(inches, lbs)
}

// TODO(wholesomeow): There's probably a better way to implement these values here
func CreateBodyType(cm float64, kg float64) npc.BodyType {
	log.Print("creating NPC body type")
	meters := cm / 100
	meters_square := meters * meters
	BMI := utilities.RoundToDecimal((kg / meters_square), 2)

	health_min := 5
	health_max := 7
	health_level := rand.Intn(health_max-health_min+1) + health_min

	body_id := makeBMI(BMI)
	body_select := health_level * body_id

	return npc.BodyType(body_select)
}

// --------------------------------------------------- CREATE NPC SEX-GENDER-SEXUAL ORIENTATION BEGIN ---------------------------------------------------
func CreateSexType() npc.SexType {
	log.Print("selecting NPC Sex")
	sex_select := rand.Intn(3) + 1
	return npc.SexType(sex_select)
}

func CreateGenderType() npc.GenderType {
	log.Print("selecting NPC Gender")
	gender_select := rand.Intn(len(npc.GenStateName)) + 1
	return npc.GenderType(gender_select)
}

// TODO(wholesomeow): Rework this to allow mixing pronouns
// TODO(wholesomeow): Rework this to be more clear with case to pronoun mapping
func CreatePronouns(gender npc.GenderType) []string {
	log.Print("selecting NPC Pronouns")
	var pronouns []string
	// TODO(wholesomeow): Rework better random selection
	r_val := rand.Intn(len(npc.Pronouns)) + 1
	switch gender {
	case 1:
		pronouns = npc.Pronouns[npc.Neutral_Pronouns]
	case 2:
		pronouns = npc.Pronouns[r_val]
	case 3: // TODO(wholesomeow): Figure out how to have sex influence pronoun selection for intersex cisgendered people
		pronouns = npc.Pronouns[r_val]
	case 4: // TODO(wholesomeow): Figure out how gender fluid people prefer to use pronouns
		pronouns = npc.Pronouns[npc.Neutral_Pronouns]
	case 5: // TODO(wholesomeow): Figure out how gender varient people prefer to use pronouns
		pronouns = npc.Pronouns[r_val]
	case 6:
		pronouns = npc.Pronouns[npc.Neutral_Pronouns]
	case 7:
		pronouns = npc.Pronouns[npc.Masc_Pronouns]
	case 8:
		pronouns = npc.Pronouns[npc.Femme_Pronouns]
	}

	return pronouns
}

func CreateOrientationType() npc.OrientationType {
	log.Print("selecting NPC Sexual Orientation")
	orientation_select := rand.Intn(len(npc.OriStateName)) + 1
	return npc.OrientationType(orientation_select)
}

// --------------------------------------------------- CREATE NPC MICE BEGIN ---------------------------------------------------
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

func remapOCEAN(value float64, minInput float64, maxInput float64, minOutput float64, maxOutput float64) float64 {
	return (value-minInput)/(maxInput-minInput)*(maxOutput-minOutput) + minOutput
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
		remapped_out := remapOCEAN(out, -200, 200, -100, 100)
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

// --------------------------------------------------- CREATE ENNEAGRAM DATA BEGIN ---------------------------------------------------
func CreateEnneagram(data EnneagramStruct, centers [][]string) Enneagram {
	log.Print("selecting NPC Enneagram")
	var enneagram Enneagram
	r_enneagram := rand.Intn(8) + 1
	enneagram.ID = r_enneagram

	// TODO(wholesomeow): Change this from random to normal distribution
	r_health := rand.Intn(8) + 1
	enneagram.LODLevel = r_health

	// Find center from correlated Enneagram selection
	for _, value := range centers {
		var num_centers []int
		split := strings.Split(string(value[2]), ",")
		for _, val := range split {
			num, err := strconv.Atoi(strings.TrimSpace(val))
			if err != nil {
				log.Fatalf("Error converting string to integer: %s", err)
			}
			num_centers = append(num_centers, num)
		}
		for idx, v := range num_centers {
			if r_enneagram == v {
				enneagram.Center = centers[idx][1]
			}
		}
	}

	// Set Dominant Emotion
	switch enneagram.Center {
	case "Thinking":
		enneagram.DominantEmotion = "Fear"
	case "Feeling":
		enneagram.DominantEmotion = "Shame"
	case "Instinctive":
		enneagram.DominantEmotion = "Anger"
	default:
		enneagram.DominantEmotion = "Default"
	}

	// Get data from selected Enneagram
	selection := data.EnneagramData[r_enneagram]

	enneagram.Archetype = selection.Archetype
	enneagram.Keywords = selection.Keywords
	enneagram.Description = selection.Description
	enneagram.Fear = selection.Fear
	enneagram.Desire = selection.Desire
	enneagram.Wings = selection.Wings
	enneagram.CurrentLOD = selection.LevelOfDevelopment[r_health]
	enneagram.LevelOfDevelopment = selection.LevelOfDevelopment
	enneagram.KeyMotivations = selection.KeyMotivations
	enneagram.Overview = selection.Overview
	enneagram.Addictions = selection.Addictions
	enneagram.GrowthRecommendations = selection.GrowthRecommendations

	return enneagram
}

func CreateOCEANText(npc_object NPCBase) string {
	log.Print("start of OCEAN Text Generation")
	var text string
	var text_slice []string

	traits := npc_object.OCEAN.Traits
	trait_name := []string{"open", "conscientious", "extraverted", "agreeable", "neurotic"}
	name := npc_object.Name
	pronouns := npc_object.Pronouns
	// TODO(wholesomeow): Come up with a better name lol
	post_pronoun := [][]string{{"is", "isn't"}, {"are", "aren't"}}
	attribute_values := []string{"not at all", "not very", "not often", "can be", "kind of", "somewhat", "sometimes", "often", "very much", "extremely"}

	log.Print("setting post pronoun selection")
	var post_pronoun_gender []string
	if pronouns[0] == "they" {
		post_pronoun_gender = post_pronoun[1]
	} else {
		post_pronoun_gender = post_pronoun[0]
	}

	// Cycle through all OCEAN values to create text
	for i := 0; i < 5; i++ {
		log.Printf("generating keyword data for trait: %s", trait_name[i])
		var attribute string
		// Determine attribute string from Aspect value
		// TODO(wholesomeow): Replace with fuzzy logic engine
		for j := 0.0; j < 100.0; j += 10.0 {
			var attribute_count float64
			if j != 0.0 {
				attribute_count = j / 10.0
			} else {
				attribute_count = 0.0
			}

			if npc_object.OCEAN.Aspect[i] < j {
				log.Printf("match found for OCEAN aspect: %s", trait_name[i])
				attribute = attribute_values[int(attribute_count)]

				// Positive or Negative post pronoun term selection
				var post_pronoun_selection string
				if attribute_count < 50 {
					post_pronoun_selection = post_pronoun_gender[0]
				} else {
					post_pronoun_selection = post_pronoun_gender[1]
				}

				// Build long and short trait descriptors
				long_traits := traits[i][:2]
				short_traits := traits[i][2:]

				long_trait_descriptor := fmt.Sprintf("%s and %s", long_traits[0], long_traits[1])
				short_trait_descriptor := fmt.Sprintf("%s, %s, and %s", short_traits[0], short_traits[1], short_traits[2])

				// Template population
				log.Printf("populating template for trait: %s", trait_name[i])
				var pro_name string
				var first_pro_name string
				if i == 0 {
					pro_name = name
					first_pro_name = "is"
				} else {
					pro_name = pronouns[0]
					first_pro_name = post_pronoun_gender[0]
				}
				// Template should read as: <Pronoun | Name> is <attribute> <Trait-Name>, which means <Pronoun> <are | aren't> <Trait-Descriptors> and <are | aren't> <Long-Trait-Descriptors>.
				template := fmt.Sprintf("%s %s %s %s, which means %s %s %s and %s %s.", pro_name, first_pro_name, attribute, trait_name[i], pronouns[0], post_pronoun_selection, long_trait_descriptor, post_pronoun_selection, short_trait_descriptor)

				text_slice = append(text_slice, template)
				break
			}
		}
	}
	text = strings.Join(text_slice, "\n")

	return text
}

// --------------------------------------------------- CREATE NPC MAIN BEGIN ---------------------------------------------------
func CreateNPC(config *configuration.Config) NPCBase {
	log.Print("start of NPC creation")
	var npc_object NPCBase
	npc_object.Name = CreateName(config)

	// Read in the CS Data csv file
	path := fmt.Sprintf("%s/%s", config.Database.CSVPath, config.Database.RequiredFiles[5])
	cognitive_data := utilities.ReadCSV(path, true)
	mice_data := cognitive_data[:4]
	cs_data := cognitive_data[4:8]
	ocean_data := cognitive_data[8:13]
	enneagram_centers := cognitive_data[13:]

	// Read in Enneagram JSON file
	path = fmt.Sprintf("%s/%s", config.Database.JSONPath, config.Database.RequiredFiles[6])
	data := utilities.ReadJSON(path)
	var enneagram_data EnneagramStruct
	err := json.Unmarshal(data, &enneagram_data)
	if err != nil {
		log.Fatalf("Failed to unmarshal json, %s", err)
	}

	// Generate Enneagram Data
	npc_object.Enneagram = CreateEnneagram(enneagram_data, enneagram_centers)

	// Generate CS and Personality Base Data
	npc_object.MICE.Aspect, npc_object.MICE.Description, npc_object.MICE.Use = CreateMICE(mice_data)
	npc_object.CS.Aspect, npc_object.CS.Data, npc_object.CS.Description, npc_object.CS.Use = CreateCSData(cs_data)
	npc_object.OCEAN.Aspect, npc_object.OCEAN.Traits, npc_object.OCEAN.Description, npc_object.OCEAN.Use = CreateOCEANData(ocean_data, npc_object.CS.Data)

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations
	log.Print("setting NPC Body Type values from Enum")
	npc_object.NPCEnums.NPCType = 0 // Set to DEFAULT on init
	npc_object.NPCType.Name = npc.NPCStateToString(npc_object.NPCEnums.NPCType)
	npc_object.NPCType.Description = npc.GetNPCStateDescription(npc_object.NPCEnums.NPCType)

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations
	ft, inch, lbs, inches := MakeSizeImperial()
	cm, kg := MakeSizeMetric(inches, lbs)
	npc_object.NPCEnums.BodyType = CreateBodyType(cm, kg)
	npc_object.BodyType.Name = npc.BodStateToString(npc_object.NPCEnums.BodyType)

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations
	log.Print("setting NPC Sex values from Enum")
	npc_object.NPCEnums.SexType = CreateSexType()
	npc_object.Sex.Name = npc.SexStateToString(npc_object.NPCEnums.SexType)

	log.Print("setting NPC Gender values from Enum")
	npc_object.NPCEnums.GenderType = CreateGenderType()
	npc_object.Gender.Name = npc.GenStateToString(npc_object.NPCEnums.GenderType)
	npc_object.Gender.Description = npc.GetGenderDescription(npc_object.NPCEnums.GenderType)

	log.Print("setting NPC Pronoun values from Enum")
	npc_object.Pronouns = CreatePronouns(npc_object.NPCEnums.GenderType)

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations
	log.Print("setting NPC Sexual Orientation values from Enum")
	npc_object.NPCEnums.OrientationType = CreateOrientationType()
	npc_object.SexualOrientation.Name = npc.OriStateToString(npc_object.NPCEnums.OrientationType)
	npc_object.SexualOrientation.Description = npc.GetOriDescription(npc_object.NPCEnums.OrientationType)

	// TOOD(wholesomeow): Create UUID function here
	log.Print("generating NPC UUID")
	npc_object.UUID = 0

	log.Print("setting NPC Appearance values")
	npc_object.NPCAppearance.Height_Ft = ft
	npc_object.NPCAppearance.Height_In = inch
	npc_object.NPCAppearance.Weight_Lbs = lbs
	npc_object.NPCAppearance.Height_Cm = cm
	npc_object.NPCAppearance.Weight_Kg = kg

	// Initial attempt at Text Generation
	log.Print("start of text generation")
	npc_object.OCEAN.Text = CreateOCEANText(npc_object)

	log.Print("NPC generation finished")
	return npc_object
}
