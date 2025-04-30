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
	// TODO(wholesomeow): Implement NPC options data for optional user-driven configuration overrides
	// Generate CS Base Data
	err = CreateCSData(&npc_object)
	if err != nil {
		return npc_object, err
	}

	// Generate REI Base Data
	err = CreateREIData(&npc_object)
	if err != nil {
		return npc_object, err
	}

	// Generate OCEAN Base Data
	err = CreateOCEANData(&npc_object)
	if err != nil {
		return npc_object, err
	}

	// Generate Enneagram Data
	err = CreateEnneagram(&npc_object)
	if err != nil {
		return npc_object, err
	}

	// Generate MICE Base Data
	err = CreateMICEData(&npc_object)
	if err != nil {
		return npc_object, err
	}

	// ----- GENERATE PHYSICALITY DATA -----
	// TODO(wholesomeow): Implement NPC options data for optional user-driven configuration overrides
	log.Print("setting NPC Type values from Enum")
	err = CreateNPCType(&npc_object)
	if err != nil {
		return npc_object, err
	}

	log.Print("setting NPC Body Type values from Enum")
	err = CreateBodyType(&npc_object)
	if err != nil {
		return npc_object, err
	}

	log.Print("setting NPC Sex values from Enum")
	err = CreateSexType(&npc_object)
	if err != nil {
		return npc_object, err
	}

	log.Print("setting NPC Gender values from Enum")
	err = CreateGenderType(&npc_object)
	if err != nil {
		return npc_object, err
	}

	log.Print("setting NPC Pronoun values from Enum")
	err = CreatePronouns(&npc_object)
	if err != nil {
		return npc_object, err
	}

	log.Print("setting NPC Sexual Orientation values from Enum")
	err = CreateOrientationType(&npc_object)
	if err != nil {
		return npc_object, err
	}

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
