package npcgen

import (
	"context"
	"log"
	"math"

	"github.com/jackc/pgx/v4"
	config "github.com/wholesomeow/npcGo/configs"
	utilities "github.com/wholesomeow/npcGo/internal/utilities"
)

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

// TODO(wholesomeow): Figure out how to use the questionare from this link https://www.psytoolkit.org/survey-library/thinking-style-rei.html
// to *actually* determine Rationality vs Experiential attributes

// TODO(wholesomeow): Figure out what I'm doing with this
func (npc_object *NPCBase) CreateREIData() error {
	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Create DB Object
	var db *pgx.Conn
	db, err = utilities.ConnectDatabase(config)
	if err != nil {
		return err
	}

	defer db.Close(context.Background())

	// Create Personality Data Query
	rei_query := "SELECT * FROM cognitive_data_npc WHERE category='REI_Dimensions'"

	// Create Personality Data Container
	rei_data, err := getREIData(db, rei_query)
	if err != nil {
		return err
	}

	cs_data := npc_object.CS.Coords

	log.Print("generating Rational-Experiential Inventory UUID")
	npc_object.REI.UUID, err = CreateUUIDv4()
	if err != nil {
		return err
	}

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

	return nil
}
