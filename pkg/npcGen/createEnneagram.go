package npcgen

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/jackc/pgx/v4"
	config "github.com/wholesomeow/npcGo/configs"
	utilities "github.com/wholesomeow/npcGo/internal/utilities"
)

func SelectEnneagram() int {
	log.Print("selecting NPC Enneagram")
	return rand.Intn(8) + 1
}

// func CreateEnneaArch(id int, data EnneagramData) string {
// 	log.Printf("populating Enneagram Archetype from selection: %d", id)
// 	return data.Archetype
// }

// func CreateEnneaCenter(id int, centers [][]string) string {
// 	log.Printf("populating Enneagram Centers from selection: %d", id)
// 	// Find center from correlated Enneagram selection
// 	var center string
// 	for _, value := range centers {
// 		var num_centers []int
// 		split := strings.Split(string(value[2]), ",")

// 		for _, val := range split {
// 			num, err := strconv.Atoi(strings.TrimSpace(val))
// 			if err != nil {
// 				log.Fatalf("Error converting string to integer: %s", err)
// 			}
// 			num_centers = append(num_centers, num)
// 		}

// 		for idx, v := range num_centers {
// 			if id == v {
// 				return centers[idx][1]
// 			}
// 		}
// 	}

// 	return center
// }

// func CreateEnneaEmote(id int, center string) string {
// 	log.Printf("populating Enneagram Dominant Emotion from selection: %d", id)
// 	// Set Dominant Emotion
// 	var emotion string
// 	switch center {
// 	case "Thinking":
// 		emotion = "Fear"
// 	case "Feeling":
// 		emotion = "Shame"
// 	case "Instinctive":
// 		emotion = "Anger"
// 	default:
// 		emotion = "Default"
// 	}

// 	return emotion
// }

// func CreateEnneaKeywords(id int, data EnneagramStruct) []string {
// 	log.Printf("populating Enneagram Keywords from selection: %d", id)
// 	return data.EnneagramData[id].Keywords
// }

// func CreateEnneaDesc(id int, data EnneagramStruct) string {
// 	log.Printf("populating Enneagram Description from selection: %d", id)
// 	return data.EnneagramData[id].Description
// }

// func CreateEnneaFear(id int, data EnneagramStruct) string {
// 	log.Printf("populating Enneagram Fear from selection: %d", id)
// 	return data.EnneagramData[id].Fear
// }

// func CreateEnneaDesire(id int, data EnneagramStruct) string {
// 	log.Printf("populating Enneagram Desire from selection: %d", id)
// 	return data.EnneagramData[id].Desire
// }

// func CreateEnneaWings(id int, data EnneagramStruct) []int {
// 	log.Printf("populating Enneagram Wings from selection: %d", id)
// 	return data.EnneagramData[id].Wings
// }

func CreateEnneaLODLevel() int {
	log.Print("selecting Enneagram Level of Health")
	// TODO(wholesomeow): Change this from random to normal distribution
	return rand.Intn(8) + 1
}

func CreateEnneaCLOD(LOD_list *[9]string, LODLevel int) string {
	log.Printf("populating Enneagram Current Level of Health from selection: %d", LODLevel)
	return LOD_list[LODLevel]
}

// func CreateEnneaLODS(id int, data EnneagramStruct) []string {
// 	log.Print("populating Enneagram Level of Health list")
// 	return data.EnneagramData[id].LevelOfDevelopment
// }

// func CreateEnneaMotive(id int, data EnneagramStruct) string {
// 	log.Printf("populating Enneagram Key Motivation from selection: %d", id)
// 	return data.EnneagramData[id].KeyMotivations
// }

// func CreateEnneaOverview(id int, data EnneagramStruct) string {
// 	log.Printf("populating Enneagram Overview from selection: %d", id)
// 	return data.EnneagramData[id].Overview
// }

// func CreateEnneaAddictions(id int, data EnneagramStruct) string {
// 	log.Printf("populating Enneagram Addictions from selection: %d", id)
// 	return data.EnneagramData[id].Addictions
// }

// func CreateEnneaGrowth(id int, data EnneagramStruct) []string {
// 	log.Printf("populating Enneagram Growth Recommendations from selection: %d", id)
// 	return data.EnneagramData[id].GrowthRecommendations
// }

func (npc_object *NPCBase) CreateEnneagram() error {
	var err error

	log.Print("generating NPC type UUID")
	npc_object.Enneagram.UUID, err = CreateUUIDv4()
	if err != nil {
		return err
	}

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

	// Create preprocess enneagram variables
	var enn_keywords string
	var enn_LOD [9]string

	// Select Enneagram here to cut down on data queried
	log.Print("selecting Enneagram number")
	var enneagram_id = SelectEnneagram()

	log.Print("querying db for Enneagram data")
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
		return err
	}

	// Populate the remaining fields
	npc_object.Enneagram.LODLevel = CreateEnneaLODLevel()
	npc_object.Enneagram.CurrentLOD = CreateEnneaCLOD(&enn_LOD, npc_object.Enneagram.LODLevel)
	npc_object.Enneagram.LevelOfDevelopment = enn_LOD

	return nil
}
