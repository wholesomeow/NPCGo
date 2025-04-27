package npcgen

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	config "github.com/wholesomeow/npcGo/configs"
	utilities "github.com/wholesomeow/npcGo/internal/utilities"
)

func GetExistingNPC(uuid string) (NPCBase, error) {
	npc_object := NPCBase{}

	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Create DB Object
	var db *pgx.Conn
	db, err = utilities.ConnectDatabase(config)
	if err != nil {
		return npc_object, err
	}

	defer db.Close(context.Background())

	// Create NPC Data Query
	npc_query := fmt.Sprintf("SELECT * FROM generated_npcs WHERE id='%s'", uuid)

	// LEFT OFF: 
	// 1. Need to create several migration files to create new tables for each
	// section of the NPC object that has a UUID and then link those tables together.
	// 2. Then, I need to query the UUIDs from the generated_npcs table into variables
	// 3. Then, query the data for those fields with those sub UUIDs
	// 4. I should probably also remove the method attribute from each of the create
	// functions and just pass in the NPC object as a parameter like before...

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data")
	err = db.QueryRow(context.Background(), npc_query).Scan(
		&npc_object.UUID,
		&npc_object.Name,
		&npc_object.Pronouns,
		&npc_object.OCEAN.UUID,
		&npc_object.OCEAN.Aspect,
		&npc_object.OCEAN.Degree,
		&npc_object.OCEAN.Traits,
		&npc_object.OCEAN.Text,
		&npc_object.OCEAN.Description,
		&npc_object.OCEAN.Use,
		&npc_object.MICE.UUID,
		&npc_object.MICE.Aspect,
		&npc_object.MICE.Degree,
		&npc_object.MICE.Traits,
		&npc_object.MICE.Text,
		&npc_object.MICE.Description,
		&npc_object.MICE.Use,
		&npc_object.CS.UUID,
		&npc_object.CS.Aspect,
		&npc_object.CS.Traits,
		&npc_object.CS.Text,
		&npc_object.CS.Coords,
		&npc_object.CS.Description,
		&npc_object.CS.Use,
		&npc_object.REI.UUID,
		&npc_object.REI.Aspect,
		&npc_object.REI.Degree,
		&npc_object.REI.Traits,
		&npc_object.REI.Text,
		&npc_object.REI.Description,
		&npc_object.REI.Use,
		&npc_object.Enneagram.UUID,
		&npc_object.Enneagram.ID,
		&npc_object.Enneagram.Archetype,
		&npc_object.Enneagram.Keywords,
		&npc_object.Enneagram.Description,
		&npc_object.Enneagram.Center,
		&npc_object.Enneagram.DominantEmotion,
		&npc_object.Enneagram.Fear,
		&npc_object.Enneagram.Desire,
		&npc_object.Enneagram.Wings,
		&npc_object.Enneagram.LODLevel,
		&npc_object.Enneagram.CurrentLOD,
		&npc_object.Enneagram.LevelOfDevelopment,
		&npc_object.Enneagram.KeyMotivations,
		&npc_object.Enneagram.Overview,
		&npc_object.Enneagram.Addictions,
		&npc_object.Enneagram.GrowthRecommendations,
		&npc_object.NPCType.UUID,
		&npc_object.NPCType.Name,
		&npc_object.NPCType.Description,
		&npc_object.NPCType.Enum,
		&npc_object.BodyType.UUID,
		&npc_object.BodyType.Name,
		&npc_object.BodyType.Enum,
		&npc_object.SexType.UUID,
		&npc_object.SexType.Name,
		&npc_object.SexType.Enum,
		&npc_object.GenderType.UUID,
		&npc_object.GenderType.Name,
		&npc_object.GenderType.Description,
		&npc_object.GenderType.Enum,
		&npc_object.SexualOrientationType.UUID,
		&npc_object.SexualOrientationType.Name,
		&npc_object.SexualOrientationType.Description,
		&npc_object.SexualOrientationType.Enum,
		&npc_object.NPCAppearance.Height_Ft,
		&npc_object.NPCAppearance.Height_In,
		&npc_object.NPCAppearance.Weight_Lbs,
		&npc_object.NPCAppearance.Height_Cm,
		&npc_object.NPCAppearance.Weight_Kg,
	)
	if err != nil {
		return npc_object, err
	}

	return npc_object, nil
}
