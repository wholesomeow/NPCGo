package npc

import (
	"go/npcGen/configuration"
	"go/npcGen/utilities"
	"log"
	"reflect"
	"testing"
)

// TestNPCCreate calls npc.CreateNPC checking for a valid return value
func TestNPCType(t *testing.T) {
	var npc NPCBase
	npc_type := reflect.TypeOf(npc)
	if npc_type != reflect.TypeOf(NPCBase{}) {
		t.Fatalf("Expected type NPCBase, got %v", reflect.TypeOf(npc))
	}
}

func TestNPCName(t *testing.T) {
	var config configuration.Config

	// Conf path hardcoded for testing purposes
	err := utilities.ChangeWorkingDir("..")
	if err != nil {
		panic(err)
	}
	conf_path := "configuration/dbconf.yaml"
	log.Printf("database conf file at path %s", conf_path)
	utilities.ReadConfig(conf_path, &config)

	test_npc := CreateNPC(&config)

	nameField := reflect.ValueOf(test_npc).FieldByName("Name")
	if nameField.Kind() != reflect.String {
		t.Fatalf("Expected Name field to be of type string, got %v", nameField.Kind())
	}
}

// TODO(wholesomeow): Implement test to gather distribution of OCEAN values over large sample size >1000
// TODO(wholesomeow): Implement test to gather distribution of CS values over large sample size >1000
