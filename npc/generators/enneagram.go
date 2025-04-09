package generators

import (
	"log"
	"math/rand"
)

// type EnneagramStruct struct {
// 	EnneagramData []struct {
// 		ID                    int      `json:"typeID"`
// 		Archetype             string   `json:"archetype"`
// 		Keywords              []string `json:"keyWords"`
// 		Description           string   `json:"briefDesc"`
// 		Fear                  string   `json:"basicFear"`
// 		Desire                string   `json:"basicDesire"`
// 		Wings                 []int    `json:"wings"`
// 		LevelOfDevelopment    []string `json:"levelOfDevelopment"`
// 		KeyMotivations        string   `json:"keyMotivations"`
// 		Overview              string   `json:"overview"`
// 		Addictions            string   `json:"addictions"`
// 		GrowthRecommendations []string `json:"growthRecommendations"`
// 	} `json:"enneagramData"`
// }

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
