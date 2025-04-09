package npc

import (
	"context"
	"fmt"
	"go/npcGen/configuration"
	"go/npcGen/database"
	npc "go/npcGen/npc/enums"
	"go/npcGen/npc/generators"
	textgen "go/npcGen/text_gen"
	"go/npcGen/utilities"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v4"
)

// --------------------------------------------------- CREATE NPC NAME BEGIN ---------------------------------------------------
func (npc *NPCBase) CreateName(config *configuration.Config) (string, error) {
	var mchain MarkovChain
	var name string
	max_attempts := 6

	log.Print("starting ngram build")
	err := mchain.BuildNGram(config, max_attempts)
	if err != nil {
		return name, err
	}

	log.Print("starting name creation")
	start_proc := time.Now()
	for count := range max_attempts {
		log.Printf("name creation attempt %d", count)
		name = mchain.MakeName()
		if mchain.CheckQuality(name) {
			break
		}
		log.Printf("name %s doesn't meet quality check... moving on to next attempt", name)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("name creation completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	return name, nil
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
	var err error
	var npc_object NPCBase

	// Create NPC Name
	npc_object.Name, err = npc_object.CreateName(config)
	if err != nil {
		return npc_object, err
	}

	// Create Personality Data Containers
	mice_data := [][]string{}
	cs_data := [][]string{}
	rei_data := [][]string{}
	ocean_data := [][]string{}

	// Create DB Object
	var db *pgx.Conn
	db, err = database.ConnectDatabase(config)
	if err != nil {
		return npc_object, err
	}

	defer db.Close(context.Background())

	// Query for required data to generate NPC
	var rows pgx.Rows
	log.Print("querying db for MICE data")
	rows, err = db.Query(context.Background(), "SELECT * FROM cognitive_data_npc WHERE category='MICE'")
	if err != nil {
		return npc_object, err
	}
	defer rows.Close()

	// Iterate through query result
	log.Print("marshalling MICE query data to slice")
	for rows.Next() {
		var id int
		var category string
		var name string
		var values string
		var description string
		var tmp []string

		err := rows.Scan(&id, &category, &name, &values, &description)
		if err != nil {
			return npc_object, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		mice_data = append(mice_data, tmp)
	}

	log.Print("querying db for CS data")
	rows, err = db.Query(context.Background(), "SELECT * FROM cognitive_data_npc WHERE category='CS_Dimensions'")
	if err != nil {
		return npc_object, err
	}
	defer rows.Close()

	// Iterate through query result
	log.Print("marshalling CS query data to slice")
	for rows.Next() {
		var id int
		var category string
		var name string
		var values string
		var description string
		var tmp []string

		err := rows.Scan(&id, &category, &name, &values, &description)
		if err != nil {
			return npc_object, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		cs_data = append(cs_data, tmp)
	}

	log.Print("querying db for REI data")
	rows, err = db.Query(context.Background(), "SELECT * FROM cognitive_data_npc WHERE category='REI_Dimensions'")
	if err != nil {
		return npc_object, err
	}
	defer rows.Close()

	// Iterate through query result
	log.Print("marshalling REI query data to slice")
	for rows.Next() {
		var id int
		var category string
		var name string
		var values string
		var description string
		var tmp []string

		err := rows.Scan(&id, &category, &name, &values, &description)
		if err != nil {
			return npc_object, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		rei_data = append(rei_data, tmp)
	}

	log.Print("querying db for OCEAN data")
	rows, err = db.Query(context.Background(), "SELECT * FROM cognitive_data_npc WHERE category='OCEAN'")
	if err != nil {
		return npc_object, err
	}
	defer rows.Close()

	// Iterate through query result
	log.Print("marshalling OCEAN query data to slice")
	for rows.Next() {
		var id int
		var category string
		var name string
		var values string
		var description string
		var tmp []string

		err := rows.Scan(&id, &category, &name, &values, &description)
		if err != nil {
			return npc_object, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		ocean_data = append(ocean_data, tmp)
	}

	// Create preprocess enneagram variables
	var enn_keywords string
	var enn_LOD [9]string

	// Select Enneagram here to cut down on data queried
	log.Print("querying db for Enneagram data")
	var enneagram_id = generators.SelectEnneagram()
	enneagram_query := fmt.Sprintf("SELECT * FROM enneagram WHERE id='%d'", enneagram_id)
	err = db.QueryRow(context.Background(), enneagram_query).Scan(
		&npc_object.Enneagram.ID,
		&npc_object.Enneagram.Archetype,
		&enn_keywords,
		&npc_object.Enneagram.Description,
		&npc_object.Enneagram.Center,
		&npc_object.Enneagram.DominantEmotion,
		&npc_object.Enneagram.Fear,
		&npc_object.Enneagram.Desire,
		&npc_object.Enneagram.Wings[0],
		&npc_object.Enneagram.Wings[1],
		&enn_LOD[0],
		&enn_LOD[1],
		&enn_LOD[2],
		&enn_LOD[3],
		&enn_LOD[4],
		&enn_LOD[5],
		&enn_LOD[6],
		&enn_LOD[7],
		&enn_LOD[8],
		&npc_object.Enneagram.KeyMotivations,
		&npc_object.Enneagram.Overview,
		&npc_object.Enneagram.Addictions,
		&npc_object.Enneagram.GrowthRecommendations[0],
		&npc_object.Enneagram.GrowthRecommendations[1],
		&npc_object.Enneagram.GrowthRecommendations[2],
		&npc_object.Enneagram.GrowthRecommendations[3],
		&npc_object.Enneagram.GrowthRecommendations[4],
	)
	if err != nil {
		return npc_object, err
	}
	defer rows.Close()

	// ----- GENERATE PERSONALITY DATA -----
	// Generate CS Base Data
	npc_object.CS.Coords = generators.CreateCSCoords()
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
	npc_object.Enneagram.LODLevel = generators.CreateEnneaLODLevel()
	npc_object.Enneagram.CurrentLOD = generators.CreateEnneaCLOD(&enn_LOD, npc_object.Enneagram.LODLevel)
	npc_object.Enneagram.LevelOfDevelopment = enn_LOD

	// Generate MICE Base Data
	mice_selection := rand.Intn(len(mice_data))
	npc_object.MICE.Aspect = generators.CreateMICEAspect(mice_selection, mice_data)
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
