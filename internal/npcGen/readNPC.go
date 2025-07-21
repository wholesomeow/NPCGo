package npcgen

import (
	"fmt"
	"log"

	config "github.com/wholesomeow/npcGo/configs"
	db "github.com/wholesomeow/npcGo/db"
)

func ReadNPC(uuids map[string]string) (NPCBase, error) {
	npc_object := NPCBase{}

	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Create DB Object
	database, err := db.ConnectDatabase(config)
	if err != nil {
		return npc_object, err
	}

	defer database.Close()

	// Create NPC Data Query
	if id, ok := uuids["npc_id"]; ok {
		npc_query := fmt.Sprintf("SELECT * FROM users.generated_npcs WHERE id='%s'", id)

		// Scan data into NPC Object
		log.Print("querying db for existing NPC data - MAIN")
		err = database.QueryRow(npc_query).Scan(
			&npc_object.UUID,
			&npc_object.Name,
			&npc_object.Pronouns,
			&npc_object.OCEAN.UUID,
			&npc_object.MICE.UUID,
			&npc_object.CS.UUID,
			&npc_object.REI.UUID,
			&npc_object.Enneagram.UUID,
			&npc_object.NPCType.UUID,
			&npc_object.BodyType.UUID,
			&npc_object.SexType.UUID,
			&npc_object.GenderType.UUID,
			&npc_object.SexualOrientationType.UUID,
			&npc_object.NPCAppearance.Height_Ft,
			&npc_object.NPCAppearance.Height_In,
			&npc_object.NPCAppearance.Weight_Lbs,
			&npc_object.NPCAppearance.Height_Cm,
			&npc_object.NPCAppearance.Weight_Kg,
		)
		if err != nil {
			return npc_object, err
		}
	}

	ocean_query := fmt.Sprintf("SELECT * FROM npc_traits.ocean_data WHERE id='%s'", npc_object.OCEAN.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - OCEAN")
	err = database.QueryRow(ocean_query).Scan(
		&npc_object.OCEAN.UUID,
		&npc_object.OCEAN.Aspect,
		&npc_object.OCEAN.Degree,
		&npc_object.OCEAN.Traits,
		&npc_object.OCEAN.Text,
		&npc_object.OCEAN.Description,
		&npc_object.OCEAN.Use,
	)
	if err != nil {
		return npc_object, err
	}

	mice_query := fmt.Sprintf("SELECT * FROM npc_traits.mice_data WHERE id='%s'", npc_object.MICE.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - MICE")
	err = database.QueryRow(mice_query).Scan(
		&npc_object.MICE.UUID,
		&npc_object.MICE.Aspect,
		&npc_object.MICE.Degree,
		&npc_object.MICE.Traits,
		&npc_object.MICE.Text,
		&npc_object.MICE.Description,
		&npc_object.MICE.Use,
	)
	if err != nil {
		return npc_object, err
	}

	cs_query := fmt.Sprintf("SELECT * FROM npc_traits.cs_data WHERE id='%s'", npc_object.CS.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - CS")
	err = database.QueryRow(cs_query).Scan(
		&npc_object.CS.UUID,
		&npc_object.CS.Aspect,
		&npc_object.CS.Traits,
		&npc_object.CS.Text,
		&npc_object.CS.Coords,
		&npc_object.CS.Description,
		&npc_object.CS.Use,
	)
	if err != nil {
		return npc_object, err
	}

	rei_query := fmt.Sprintf("SELECT * FROM npc_traits.rei_data WHERE id='%s'", npc_object.REI.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - REI")
	err = database.QueryRow(rei_query).Scan(
		&npc_object.REI.UUID,
		&npc_object.REI.Aspect,
		&npc_object.REI.Degree,
		&npc_object.REI.Traits,
		&npc_object.REI.Text,
		&npc_object.REI.Description,
		&npc_object.REI.Use,
	)
	if err != nil {
		return npc_object, err
	}

	enn_query := fmt.Sprintf("SELECT * FROM npc_traits.enneagram_data WHERE id='%s'", npc_object.Enneagram.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - Enneagram")
	err = database.QueryRow(enn_query).Scan(
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
	)
	if err != nil {
		return npc_object, err
	}

	type_query := fmt.Sprintf("SELECT * FROM npc_meta.npc_types WHERE id='%s'", npc_object.NPCType.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - NPC Type")
	err = database.QueryRow(type_query).Scan(
		&npc_object.NPCType.UUID,
		&npc_object.NPCType.Name,
		&npc_object.NPCType.Description,
		&npc_object.NPCType.Enum,
	)
	if err != nil {
		return npc_object, err
	}

	body_query := fmt.Sprintf("SELECT * FROM npc_meta.body_types WHERE id='%s'", npc_object.BodyType.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - NPC Body")
	err = database.QueryRow(body_query).Scan(
		&npc_object.BodyType.UUID,
		&npc_object.BodyType.Name,
		&npc_object.BodyType.Enum,
	)
	if err != nil {
		return npc_object, err
	}

	sex_query := fmt.Sprintf("SELECT * FROM npc_meta.sex_types WHERE id='%s'", npc_object.SexType.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - NPC Sex")
	err = database.QueryRow(sex_query).Scan(
		&npc_object.SexType.UUID,
		&npc_object.SexType.Name,
		&npc_object.SexType.Enum,
	)
	if err != nil {
		return npc_object, err
	}

	gender_query := fmt.Sprintf("SELECT * FROM npc_meta.gender_types WHERE id='%s'", npc_object.GenderType.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - NPC Gender")
	err = database.QueryRow(gender_query).Scan(
		&npc_object.GenderType.UUID,
		&npc_object.GenderType.Name,
		&npc_object.GenderType.Description,
		&npc_object.GenderType.Enum,
	)
	if err != nil {
		return npc_object, err
	}

	ori_query := fmt.Sprintf("SELECT * FROM npc_meta.orientation_types WHERE id='%s'", npc_object.SexualOrientationType.UUID)

	// Scan data into NPC Object
	log.Print("querying db for existing NPC data - NPC Orientation")
	err = database.QueryRow(ori_query).Scan(
		&npc_object.SexualOrientationType.UUID,
		&npc_object.SexualOrientationType.Name,
		&npc_object.SexualOrientationType.Description,
		&npc_object.SexualOrientationType.Enum,
	)
	if err != nil {
		return npc_object, err
	}

	return npc_object, nil
}
