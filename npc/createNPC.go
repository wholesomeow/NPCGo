package npc

import (
	"go/npcGen/configuration"
	npc "go/npcGen/npc/enums"
	"go/npcGen/utilities"
	"log"
	"math/rand"
	"time"
)

func CreateName(config *configuration.Config) string {
	var mchain MarkovChain
	var name string
	max_attempts := 6

	buildNGram(&mchain, config, max_attempts)

	log.Print("starting name creation")
	start_proc := time.Now()
	for count := range max_attempts {
		log.Printf("name creation attempt %d", count)
		name = makeName(&mchain)
		if checkQuality(&mchain, name) {
			break
		}
		log.Printf("name %s doesn't meet quality check... moving on to next attempt", name)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("name creation completed... elapsed time: %s", time.Duration.String(elapsed_proc))

	return name
}

// TODO(wholesomeow): There's probably a better way to implement these values here
func makeBMI(BMI float64) int {
	if BMI <= 18.5 {
		return 1
	} else if 18.5 < BMI || BMI <= 24.9 {
		return 2
	} else if 25 < BMI || BMI <= 29.29 {
		return 3
	} else {
		return 4
	}
}

// TODO(wholesomeow): There's probably a better way to implement these values here
func MakeSizeImperial() (int, int, int, int) {
	ft := 0
	inch := 0
	lbs := 0

	ft_medium := []int{4, 5, 6, 7}
	ft_small := []int{2, 3}

	lbs_min := 110
	lbs_max := 250

	r_height := rand.Intn(2)
	if r_height == 0 {
		ft = ft_medium[rand.Intn(len(ft_medium))]
		inch = rand.Intn(11)
		lbs = rand.Intn(lbs_max-lbs_min+1) + lbs_min
	} else {
		ft = ft_small[rand.Intn(len(ft_small))]
		inch = rand.Intn(11)
		lbs = rand.Intn(lbs_max-lbs_min+1) + lbs_min
	}

	inches := (ft * 12) + inch

	return ft, inch, lbs, inches
}

func MakeSizeMetric(inches int, lbs int) (float64, float64) {
	return utilities.ImperialToMetric(inches, lbs)
}

// TODO(wholesomeow): There's probably a better way to implement these values here
func CreateBodyType(cm float64, kg float64) npc.BodyType {
	meters := cm / 100
	meters_square := meters * meters
	BMI := utilities.RoundToDecimal((kg / meters_square), 2)

	health_min := 5
	health_max := 7
	health_level := rand.Intn(health_max-health_min+1) + health_min

	body_id := makeBMI(BMI)
	body_select := health_level * body_id

	return npc.BodyType(body_select)
}

func CreateGenderType() npc.GenderType {
	gender_select := rand.Intn(8) + 1
	return npc.GenderType(gender_select)
}

// TODO(wholesomeow): Rework this to allow with mixing pronouns
func CreatePronouns(gender npc.GenderType) string {
	var pronouns string
	r_val := rand.Intn(3) + 1
	switch gender {
	case 1:
		pronouns = npc.Pronouns[4][0]
	case 2:
		pronouns = npc.Pronouns[r_val][0]
	case 3: // TODO(wholesomeow): Figure out how to have sex influence pronoun selection for intersex cisgendered people
		pronouns = npc.Pronouns[r_val][0]
	case 4: // TODO(wholesomeow): Figure out how gender fluid people prefer to use pronouns
		pronouns = npc.Pronouns[3][0]
	case 5: // TODO(wholesomeow): Figure out how gender varient people prefer to use pronouns
		pronouns = npc.Pronouns[r_val][0]
	case 6:
		pronouns = npc.Pronouns[4][0]
	case 7:
		pronouns = npc.Pronouns[1][0]
	case 8:
		pronouns = npc.Pronouns[2][0]
	}

	return pronouns
}

func CreateNPC(config *configuration.Config) NPCBase {
	var npc NPCBase
	npc.Name = CreateName(config)

	// TODO(wholesomeow): Implement enums into NPC here
	npc.NPCEnums.NPCType = 0 // Set to DEFAULT on init

	ft, inch, lbs, inches := MakeSizeImperial()
	cm, kg := MakeSizeMetric(inches, lbs)
	npc.NPCEnums.BodyType = CreateBodyType(cm, kg)

	npc.NPCEnums.GenderType = CreateGenderType()
	npc.Pronouns = CreatePronouns(npc.NPCEnums.GenderType)

	npc.NPCAppearance.Height_Ft = ft
	npc.NPCAppearance.Height_In = inch
	npc.NPCAppearance.Weight_Lbs = lbs
	npc.NPCAppearance.Height_Cm = cm
	npc.NPCAppearance.Weight_Kg = kg

	return npc
}
