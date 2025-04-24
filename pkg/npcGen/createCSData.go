package npcgen

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	config "github.com/wholesomeow/npcGo/configs"
	utilities "github.com/wholesomeow/npcGo/internal/utilities"
)

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

func coordsToSelection(cs_coords [2]int) int {
	var selection int
	if cs_coords[0] <= 0 && cs_coords[1] <= 0 {
		selection = 0
	} else if cs_coords[0] <= 0 && cs_coords[1] >= 0 {
		selection = 1
	} else if cs_coords[0] >= 0 && cs_coords[1] >= 0 {
		selection = 2
	} else if cs_coords[0] >= 0 && cs_coords[1] <= 0 {
		selection = 3
	}

	return selection
}

func (npc_object *NPCBase) CreateCSData() error {
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

	// Create Personality Data Queries
	cs_query := "SELECT * FROM cognitive_data_npc WHERE category='CS_Dimensions'"

	// Create Personality Data Containers
	cs_data, err := getCSData(db, cs_query)
	if err != nil {
		return err
	}

	log.Print("generating Cognitive Science UUID")
	npc_object.CS.UUID, err = CreateUUIDv4()
	if err != nil {
		return err
	}

	log.Print("generating Cognitive Science Aspect for NPC")
	idx := coordsToSelection(npc_object.CS.Coords)
	selection := cs_data[idx]
	npc_object.CS.Aspect = selection[1]

	log.Print("generating Cognitive Science Coordiantes for NPC")
	npc_object.CS.Coords[0] = utilities.RandomRange(-100, 100)
	npc_object.CS.Coords[1] = utilities.RandomRange(-100, 100)

	// TODO(wholesomeow): Create the logic for this
	log.Print("generating Cognitive Science Traits for NPC")
	traits := []string{}
	npc_object.CS.Traits = traits

	log.Print("populating Cognitive Science Description for NPC")
	npc_object.CS.Description = selection[2]

	log.Print("populating Cognitive Science Use for NPC")
	npc_object.CS.Use = "used to quantify at which cognitive aspects a person either excels at, struggles with, or a combination of both"

	return nil
}
