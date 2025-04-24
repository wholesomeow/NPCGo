package npcgen

import (
	"log"

	namegen "github.com/wholesomeow/npcGo/pkg/nameGen"
	textgen "github.com/wholesomeow/npcGo/pkg/textGen"
)

func CreateNPC() (NPCBase, error) {
	log.Print("start of NPC creation")
	npc_object := NPCBase{}
	var err error

	log.Print("generating NPC UUID")
	npc_object.UUID, err = CreateUUIDv4()
	if err != nil {
		return npc_object, err
	}

	// Create NPC Name
	npc_object.Name, err = namegen.CreateName()
	if err != nil {
		return npc_object, err
	}

	// ----- GENERATE PERSONALITY DATA -----
	// Generate CS Base Data
	npc_object.CreateCSData()

	// Generate REI Base Data
	npc_object.CreateREIData()

	// Generate OCEAN Base Data
	npc_object.CreateOCEANData()

	// Generate Enneagram Data
	npc_object.CreateEnneagram()

	// Generate MICE Base Data
	npc_object.CreateMICEData()

	// ----- GENERATE PHYSICALITY DATA -----
	// TODO(wholesomeow): Implement NPC options data for optional user-driven configuration overrides
	log.Print("setting NPC Type values from Enum")
	npc_object.NPCType.Enum = 0 // Set to DEFAULT on init
	npc_object.NPCType.Name = npc_object.NPCType.Enum.NPCStateToString()
	npc_object.NPCType.Description = npc_object.NPCType.Enum.GetNPCStateDescription()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Body Type values from Enum")
	npc_object.MakeSizeImperial()
	npc_object.MakeSizeMetric()
	npc_object.CreateBodyType()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Sex values from Enum")
	npc_object.CreateSexType()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Gender values from Enum")
	npc_object.CreateGenderType()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Pronoun values from Enum")
	npc_object.CreatePronouns()

	// TODO(wholesomeow): Implement NPC options data for optional user-driven configurations overrides
	log.Print("setting NPC Sexual Orientation values from Enum")
	npc_object.CreateOrientationType()

	// ----- GENERATE TEXT -----
	log.Print("start of text generation")
	OCEANTextData := CreateOCEANText(npc_object.Name,
		npc_object.Pronouns,
		npc_object.OCEAN.Traits,
		npc_object.OCEAN.Degree,
	)
	npc_object.OCEAN.Text = textgen.SimpleSentenceBuilder(OCEANTextData)

	log.Print("NPC generation finished")
	return npc_object, nil
}
