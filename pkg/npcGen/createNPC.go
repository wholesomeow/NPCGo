package npcgen

import (
	"context"
	"fmt"
	config "go/npcGen/configs"
	"go/npcGen/internal/utilities"
	namegen "go/npcGen/pkg/nameGen"
	textgen "go/npcGen/pkg/textGen"
	"log"
	"math/rand"

	"github.com/jackc/pgx/v4"
)

func getMICEData(db *pgx.Conn, q_str string) ([][]string, error) {
	data := [][]string{}

	// Query for required data to generate NPC
	var rows pgx.Rows
	log.Print("querying db for MICE data")
	rows, err := db.Query(context.Background(), q_str)
	if err != nil {
		return data, err
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
			return data, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		data = append(data, tmp)
	}
	defer rows.Close()

	return data, nil
}

func getCSData(db *pgx.Conn, q_str string) ([][]string, error) {
	data := [][]string{}

	// Query for required data to generate NPC
	log.Print("querying db for CS data")
	rows, err := db.Query(context.Background(), q_str)
	if err != nil {
		return data, err
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
			return data, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		data = append(data, tmp)
	}
	defer rows.Close()

	return data, nil
}

func getREIData(db *pgx.Conn, q_str string) ([][]string, error) {
	data := [][]string{}

	// Query for required data to generate NPC
	log.Print("querying db for REI data")
	rows, err := db.Query(context.Background(), q_str)
	if err != nil {
		return data, err
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
			return data, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		data = append(data, tmp)
	}
	defer rows.Close()

	return data, err
}

func getOCEANData(db *pgx.Conn, q_str string) ([][]string, error) {
	data := [][]string{}

	// Query for required data to generate NPC
	log.Print("querying db for OCEAN data")
	rows, err := db.Query(context.Background(), q_str)
	if err != nil {
		return data, err
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
			return data, err
		}

		tmp = append(tmp, name)
		tmp = append(tmp, values)
		tmp = append(tmp, description)

		data = append(data, tmp)
	}
	defer rows.Close()

	return data, err
}

// --------------------------------------------------- CREATE NPC MAIN BEGIN ---------------------------------------------------
func CreateNPC(config *config.Config) (NPCBase, error) {
	log.Print("start of NPC creation")
	npc_object := NPCBase{}

	// Create DB Object
	var db *pgx.Conn
	db, err := utilities.ConnectDatabase(config)
	if err != nil {
		return npc_object, err
	}

	defer db.Close(context.Background())

	// Create Personality Data Queries
	mice_query := "SELECT * FROM cognitive_data_npc WHERE category='MICE'"
	cs_query := "SELECT * FROM cognitive_data_npc WHERE category='CS_Dimensions'"
	rei_query := "SELECT * FROM cognitive_data_npc WHERE category='REI_Dimensions'"
	ocean_query := "SELECT * FROM cognitive_data_npc WHERE category='OCEAN'"

	// Create Personality Data Containers
	mice_data, err := getMICEData(db, mice_query)
	if err != nil {
		return npc_object, err
	}

	cs_data, err := getCSData(db, cs_query)
	if err != nil {
		return npc_object, err
	}

	rei_data, err := getREIData(db, rei_query)
	if err != nil {
		return npc_object, err
	}

	ocean_data, err := getOCEANData(db, ocean_query)
	if err != nil {
		return npc_object, err
	}

	// Create preprocess enneagram variables
	var enn_keywords string
	var enn_LOD [9]string

	// Select Enneagram here to cut down on data queried
	log.Print("querying db for Enneagram data")
	var enneagram_id = SelectEnneagram()
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

	// Create NPC Name
	npc_object.Name, err = namegen.CreateName(config)
	if err != nil {
		return npc_object, err
	}

	// ----- GENERATE PERSONALITY DATA -----
	// Generate CS Base Data
	npc_object.CreateCSData(cs_data)

	// Generate REI Base Data
	npc_object.CreateREIData(rei_data)

	// Generate OCEAN Base Data
	npc_object.CreateOCEANData(ocean_data)

	// Generate Enneagram Data
	npc_object.Enneagram.LODLevel = CreateEnneaLODLevel()
	npc_object.Enneagram.CurrentLOD = CreateEnneaCLOD(&enn_LOD, npc_object.Enneagram.LODLevel)
	npc_object.Enneagram.LevelOfDevelopment = enn_LOD

	// Generate MICE Base Data
	mice_selection := rand.Intn(len(mice_data))
	npc_object.CreateMICEData(mice_selection, mice_data)

	// ----- GENERATE PHYSICALITY DATA -----
	// TODO(wholesomeow): Implement NPC options data for optional user-driven configuration overrides
	log.Print("setting NPC Type values from Enum")
	npc_object.NPCType.Enum = 0 // Set to DEFAULT on init
	npc_object.NPCType.Name = npc_object.NPCType.Enum.NPCStateToString()
	npc_object.NPCType.Description = npc_object.NPCType.Enum.GetNPCStateDescription()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Body Type values from Enum")
	npc_object.MakeSizeImperial()
	npc_object.MakeSizeMetric()
	npc_object.CreateBodyType()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Sex values from Enum")
	npc_object.CreateSexType()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Gender values from Enum")
	npc_object.CreateGenderType()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Pronoun values from Enum")
	npc_object.CreatePronouns()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Sexual Orientation values from Enum")
	npc_object.CreateOrientationType()

	// TOOD(wholesomeow): Create UUID function here
	log.Print("generating NPC UUID")
	npc_object.UUID = 0

	// ----- GENERATE TEXT -----
	log.Print("start of text generation")
	OCEANTextData := CreateOCEANText(npc_object.Name,
		npc_object.Pronouns,
		npc_object.OCEAN.Traits,
		npc_object.OCEAN.Degree,
	)
	npc_object.OCEAN.Text = textgen.SimpleSentenceBuilder(OCEANTextData)

	log.Print("NPC generation finished")
	return npc_object, nil
}
