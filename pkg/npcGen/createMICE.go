package npcgen

import (
	"context"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4"
	config "github.com/wholesomeow/npcGo/configs"
	utilities "github.com/wholesomeow/npcGo/internal/utilities"
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

func (npc_object *NPCBase) CreateMICEData() error {
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
	mice_query := "SELECT * FROM cognitive_data_npc WHERE category='MICE'"

	// Create Personality Data Container
	mice_data, err := getMICEData(db, mice_query)
	if err != nil {
		return err
	}

	r_val := rand.Intn(len(mice_data))
	cs_data := npc_object.CS.Coords

	log.Print("generating MICE UUID")
	npc_object.MICE.UUID, err = CreateUUIDv4()
	if err != nil {
		return err
	}

	log.Print("setting MICE values for NPC")
	selection := mice_data[r_val]
	npc_object.MICE.Aspect = selection[1]

	log.Print("creating MICE Degree value")
	mice_cast := []float64{}
	// X Coord cast first
	split := strings.Split(string(selection[1]), ",")
	x, err := strconv.Atoi(strings.TrimSpace(split[0]))
	if err != nil {
		log.Fatalf("Error converting string to X coordinate integer: %s", err)
	}
	mice_cast = append(mice_cast, float64(x))

	// Y Coord cast second
	y, err := strconv.Atoi(strings.TrimSpace(split[1]))
	if err != nil {
		log.Fatalf("Error converting string to Y coordinate integer: %s", err)
	}
	mice_cast = append(mice_cast, float64(y))

	// Variable casting
	x1 := mice_cast[0]
	y1 := mice_cast[1]
	x2 := float64(cs_data[0])
	y2 := float64(cs_data[1])

	out := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	npc_object.MICE.Degree = utilities.RemapInt(out, -250, 250, -100, 100)

	// TODO(wholesomeow): Create the logic for this
	// TODO(wholesomeow): Figure out how I'm going to create a traits list to describe
	// how someone could be convinced/manipulated
	log.Print("creating MICE Traits")
	traits := []string{}
	npc_object.MICE.Traits = traits

	log.Print("setting MICE Description")
	log.Print("selecting specifc MICE description at index: 3")
	npc_object.MICE.Description = selection[2]

	log.Print("setting MICE Usage")
	npc_object.MICE.Use = "used to list the primary reasons why someone would become a spy, insider threat, or collaborate with a hostile organization"

	return nil
}
