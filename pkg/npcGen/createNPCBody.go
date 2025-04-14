package npcgen

import (
	"go/npcGen/internal/utilities"
	"go/npcGen/pkg/npcGen/enums"
	"log"
	"math/rand"
)

func makeBMI(BMI float64) int {
	log.Print("creating NPC BMI")
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

func (npc_object *NPCBase) MakeSizeImperial() {
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

	log.Print("setting NPC Appearance values")
	npc_object.NPCAppearance.Height_Ft = ft
	npc_object.NPCAppearance.Height_In = inches
	npc_object.NPCAppearance.Weight_Lbs = lbs
}

func imperialToMetric(inches int, lbs int) (float64, float64) {
	cm := float64(inches) * 2.54
	kg := float64(lbs) * 0.453592

	return cm, kg
}

func (npc_object *NPCBase) MakeSizeMetric() {
	inches := npc_object.NPCAppearance.Height_In
	lbs := npc_object.NPCAppearance.Weight_Lbs

	cm, kg := imperialToMetric(inches, lbs)

	npc_object.NPCAppearance.Height_Cm = cm
	npc_object.NPCAppearance.Weight_Kg = kg
}

func (npc_object *NPCBase) CreateBodyType() {
	log.Print("creating NPC body type")
	meters := npc_object.NPCAppearance.Height_Cm / 100
	meters_square := meters * meters
	BMI := utilities.RoundToDecimal((npc_object.NPCAppearance.Weight_Kg / meters_square), 2)

	health_min := 5
	health_max := 7
	health_level := rand.Intn(health_max-health_min+1) + health_min

	body_id := makeBMI(BMI)
	body_select := health_level * body_id

	npc_object.BodyType.Enum = enums.BodyType(body_select)
	npc_object.BodyType.Name = npc_object.BodyType.Enum.BodStateToString()
}
