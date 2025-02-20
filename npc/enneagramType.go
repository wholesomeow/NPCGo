package npc

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
