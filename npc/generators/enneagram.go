package generators

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type EnneagramStruct struct {
	EnneagramData []struct {
		ID                    int      `json:"typeID"`
		Archetype             string   `json:"archetype"`
		Keywords              []string `json:"keyWords"`
		Description           string   `json:"briefDesc"`
		Fear                  string   `json:"basicFear"`
		Desire                string   `json:"basicDesire"`
		Wings                 []int    `json:"wings"`
		LevelOfDevelopment    []string `json:"levelOfDevelopment"`
		KeyMotivations        string   `json:"keyMotivations"`
		Overview              string   `json:"overview"`
		Addictions            string   `json:"addictions"`
		GrowthRecommendations []string `json:"growthRecommendations"`
	} `json:"enneagramData"`
}

type Enneagram struct {
	ID                    int
	Archetype             string
	Center                string
	DominantEmotion       string
	Keywords              []string
	Description           string
	Fear                  string
	Desire                string
	Wings                 []int
	LODLevel              int
	CurrentLOD            string
	LevelOfDevelopment    []string
	KeyMotivations        string
	Overview              string
	Addictions            string
	GrowthRecommendations []string
}

func CreateEnneagram(data EnneagramStruct, centers [][]string) Enneagram {
	log.Print("selecting NPC Enneagram")
	var enneagram Enneagram
	r_enneagram := rand.Intn(8) + 1
	enneagram.ID = r_enneagram

	// TODO(wholesomeow): Change this from random to normal distribution
	r_health := rand.Intn(8) + 1
	enneagram.LODLevel = r_health

	// Find center from correlated Enneagram selection
	for _, value := range centers {
		var num_centers []int
		split := strings.Split(string(value[2]), ",")
		for _, val := range split {
			num, err := strconv.Atoi(strings.TrimSpace(val))
			if err != nil {
				log.Fatalf("Error converting string to integer: %s", err)
			}
			num_centers = append(num_centers, num)
		}
		for idx, v := range num_centers {
			if r_enneagram == v {
				enneagram.Center = centers[idx][1]
			}
		}
	}

	// Set Dominant Emotion
	switch enneagram.Center {
	case "Thinking":
		enneagram.DominantEmotion = "Fear"
	case "Feeling":
		enneagram.DominantEmotion = "Shame"
	case "Instinctive":
		enneagram.DominantEmotion = "Anger"
	default:
		enneagram.DominantEmotion = "Default"
	}

	// Get data from selected Enneagram
	selection := data.EnneagramData[r_enneagram]

	enneagram.Archetype = selection.Archetype
	enneagram.Keywords = selection.Keywords
	enneagram.Description = selection.Description
	enneagram.Fear = selection.Fear
	enneagram.Desire = selection.Desire
	enneagram.Wings = selection.Wings
	enneagram.CurrentLOD = selection.LevelOfDevelopment[r_health]
	enneagram.LevelOfDevelopment = selection.LevelOfDevelopment
	enneagram.KeyMotivations = selection.KeyMotivations
	enneagram.Overview = selection.Overview
	enneagram.Addictions = selection.Addictions
	enneagram.GrowthRecommendations = selection.GrowthRecommendations

	return enneagram
}
