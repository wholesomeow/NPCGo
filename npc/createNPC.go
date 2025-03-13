package npc

import (
	"encoding/json"
	"fmt"
	"go/npcGen/configuration"
	npc "go/npcGen/npc/enums"
	"go/npcGen/npc/generators"
	textgen "go/npcGen/text_gen"
	"go/npcGen/utilities"
	"log"
	"math/rand"
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

// --------------------------------------------------- CREATE NPC MAIN BEGIN ---------------------------------------------------
func CreateNPC(config *configuration.Config) (NPCBase, error) {
	log.Print("start of NPC creation")
	var npc_object NPCBase
	npc_object.Name = CreateName(config)

	//Create Personality Data Containers
	mice_data := [][]string{}
	cs_data := [][]string{}
	rei_data := [][]string{}
	ocean_data := [][]string{}
	enneagram_centers := [][]string{}
	var enneagram_data generators.EnneagramStruct

	mode := strings.ToLower(config.Server.Mode)
	if mode == "dev" {
		log.Print("read in all required NPC data from files")
		// Read in the CS Data csv file
		// TODO(wholesomeow): I don't like the hardcoded file name here, need to fix
		path := fmt.Sprintf("%s/%s", config.Database.CSVPath, "NPC_Cognitive_Data.csv")
		cognitive_data, err := utilities.ReadCSV(path, true)
		if err != nil {
			return npc_object, err
		}
		mice_data = cognitive_data[:4]
		cs_data = cognitive_data[4:8]
		rei_data = cognitive_data[8:12]
		ocean_data = cognitive_data[12:17]
		enneagram_centers = cognitive_data[17:]

		// Read in Enneagram JSON file
		// TODO(wholesomeow): I don't like the hardcoded file name here, need to fix
		path = fmt.Sprintf("%s/%s", config.Database.JSONPath, "enneagramData.json")
		data, err := utilities.ReadJSON(path)
		if err != nil {
			log.Fatalf("Failed to read json, %s", err)
		}
		err = json.Unmarshal(data, &enneagram_data)
		if err != nil {
			return npc_object, err
		}
	} else {
		log.Print("establishing connection to database")
	}

	// ----- GENERATE PERSONALITY DATA -----
	// Generate CS Base Data
	npc_object.CS.Coords = generators.CreateCSCoords(cs_data)
	npc_object.CS.Aspect = generators.CreateCSAspect(cs_data, npc_object.CS.Coords)
	// TODO(wholesomeow): Create the logic for this
	npc_object.CS.Traits = generators.CreateCSTraits(cs_data, npc_object.CS.Coords)
	npc_object.CS.Description = generators.CreateCSDesc(cs_data, npc_object.CS.Coords)
	npc_object.CS.Use = generators.CreateCSUse()

	// Generate REI Base Data
	npc_object.REI.Aspect = generators.CreateREIAspect(npc_object.CS.Coords)
	npc_object.REI.Degree = generators.CreateREIDegree(npc_object.CS.Coords)
	// TODO(wholesomeow): Create the logic for this
	npc_object.REI.Traits = generators.CreateREITraits()
	npc_object.REI.Description = generators.CreateREIDesc(rei_data, npc_object.CS.Coords)
	npc_object.REI.Use = generators.CreateREIUse()

	// Generate OCEAN Base Data
	npc_object.OCEAN.Aspect = generators.CreateOCEANAspect(ocean_data, npc_object.CS.Coords)
	npc_object.OCEAN.Degree = generators.CreateOCEANDegree(ocean_data, npc_object.CS.Coords)
	npc_object.OCEAN.Traits = generators.CreateOCEANTraits()
	npc_object.OCEAN.Description = generators.CreateOCEANDesc()
	npc_object.OCEAN.Use = generators.CreateOCEANUse()

	// Generate Enneagram Data
	npc_object.Enneagram.ID = generators.SelectEnneagram()
	npc_object.Enneagram.Archetype = generators.CreateEnneaArch(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.Center = generators.CreateEnneaCenter(npc_object.Enneagram.ID, enneagram_centers)
	npc_object.Enneagram.DominantEmotion = generators.CreateEnneaEmote(npc_object.Enneagram.ID,
		npc_object.Enneagram.Center,
	)
	npc_object.Enneagram.Keywords = generators.CreateEnneaKeywords(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.Description = generators.CreateEnneaDesc(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.Fear = generators.CreateEnneaFear(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.Desire = generators.CreateEnneaDesire(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.Wings = generators.CreateEnneaWings(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.LODLevel = generators.CreateEnneaLODLevel()
	npc_object.Enneagram.CurrentLOD = generators.CreateEnneaCLOD(npc_object.Enneagram.ID,
		npc_object.Enneagram.LODLevel,
		enneagram_data,
	)
	npc_object.Enneagram.LevelOfDevelopment = generators.CreateEnneaLODS(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.KeyMotivations = generators.CreateEnneaMotive(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.Overview = generators.CreateEnneaOverview(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.Addictions = generators.CreateEnneaAddictions(npc_object.Enneagram.ID, enneagram_data)
	npc_object.Enneagram.GrowthRecommendations = generators.CreateEnneaGrowth(npc_object.Enneagram.ID, enneagram_data)

	// Generate MICE Base Data
	mice_selection := rand.Intn(len(mice_data))
	npc_object.MICE.Aspect = generators.CreateMICEAspect(mice_selection, mice_data, npc_object.CS.Coords)
	npc_object.MICE.Degree = generators.CreateMICEDegree(mice_selection, mice_data, npc_object.CS.Coords)
	// TODO(wholesomeow): Create the logic for this
	npc_object.MICE.Traits = generators.CreateMICETraits(mice_selection, mice_data, npc_object.CS.Coords)
	npc_object.MICE.Description = generators.CreateMICEDesc(mice_selection, mice_data, npc_object.CS.Coords)
	npc_object.MICE.Use = generators.CreateMICEUse()

	// ----- GENERATE PHYSICALITY DATA -----
	// TODO(wholesomeow): Implement NPC options data for optional user-driven configuration overrides
	log.Print("setting NPC Body Type values from Enum")
	npc_object.NPCEnums.NPCType = 0 // Set to DEFAULT on init
	npc_object.NPCType.Name = npc.NPCStateToString(npc_object.NPCEnums.NPCType)
	npc_object.NPCType.Description = npc.GetNPCStateDescription(npc_object.NPCEnums.NPCType)

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	ft, inch, lbs, inches := MakeSizeImperial()
	cm, kg := MakeSizeMetric(inches, lbs)
	npc_object.NPCEnums.BodyType = CreateBodyType(cm, kg)
	npc_object.BodyType.Name = npc.BodStateToString(npc_object.NPCEnums.BodyType)

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Sex values from Enum")
	npc_object.NPCEnums.SexType = CreateSexType()
	npc_object.Sex.Name = npc.SexStateToString(npc_object.NPCEnums.SexType)

	log.Print("setting NPC Gender values from Enum")
	npc_object.NPCEnums.GenderType = CreateGenderType()
	npc_object.Gender.Name = npc.GenStateToString(npc_object.NPCEnums.GenderType)
	npc_object.Gender.Description = npc.GetGenderDescription(npc_object.NPCEnums.GenderType)

	log.Print("setting NPC Pronoun values from Enum")
	npc_object.Pronouns = CreatePronouns(npc_object.NPCEnums.GenderType)

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
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

	// ----- GENERATE TEXT -----
	log.Print("start of text generation")
	OCEANTextData := generators.CreateOCEANText(npc_object.Name,
		npc_object.Pronouns,
		npc_object.OCEAN.Traits,
		npc_object.OCEAN.Degree,
	)
	npc_object.OCEAN.Text = textgen.SimpleSentenceBuilder(OCEANTextData)

	log.Print("NPC generation finished")
	return npc_object, nil
}
