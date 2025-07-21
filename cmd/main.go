package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	config "github.com/wholesomeow/npcGo/configs"
	db "github.com/wholesomeow/npcGo/db"
	npcapi "github.com/wholesomeow/npcGo/internal/api"
	npcapi_npc "github.com/wholesomeow/npcGo/internal/api/npc"
	npcapi_reroll_enum "github.com/wholesomeow/npcGo/internal/api/reroll/enum"
	npcapi_reroll_npc "github.com/wholesomeow/npcGo/internal/api/reroll/npc"
)

func startAPIServer() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(npcapi.Response404)

	// TODO(wholesomeow): User related endpoints

	// TODO(wholesomeow): Admin realted endpoints? Maybe?

	// Main NPC API group
	npc_group := router.Group("/npc")
	npc_group.POST("/generate", npcapi_npc.APICreateNPC)

	// TODO(wholesomeow): NPC info collection endpoints maybe?
	// I'm not sure what exactly, maybe NPC save and load?

	// Sub NPC API group - Reroll NPC values
	npc_reroll := npc_group.Group("/reroll")
	npc_reroll.GET("/name", npcapi_reroll_npc.APICreateName)
	npc_reroll.GET("/cs", npcapi_reroll_npc.APICreateCS)
	npc_reroll.GET("/ocean", npcapi_reroll_npc.APICreateOCEAN)
	npc_reroll.GET("/mice", npcapi_reroll_npc.APICreateMICE)
	npc_reroll.GET("/rei", npcapi_reroll_npc.APICreateREI)
	npc_reroll.GET("/enneagram", npcapi_reroll_npc.APICreateEnneagram)

	// Sub NPC API group - Reroll NPC Enum values
	enum_reroll := npc_reroll.Group("/enum")
	enum_reroll.GET("/bodyType", npcapi_reroll_enum.APICreateBody)
	enum_reroll.GET("/genderType", npcapi_reroll_enum.APICreateGender)
	enum_reroll.GET("/npcType", npcapi_reroll_enum.APICreateType)
	enum_reroll.GET("/orientationType", npcapi_reroll_enum.APICreateSexOri)
	enum_reroll.GET("/sexType", npcapi_reroll_enum.APICreateSex)

	router.Run("0.0.0.0:8080")
}

func printHelp() {
	helpText := `
				Usage:
				npcgen <command>

				Available Commands:
				server            Start the NPC API server
				init              Initialize the database (create tables, schemas, etc.)
				version           Print the current version of npcgen
				help              Show this help message

				Examples:
				npcgen server
				npcgen init
				npcgen version

				Use "npcgen help" for more information about a command.
				`

	fmt.Fprintln(os.Stderr, helpText)
}

func main() {
	log.Print("Staring NPCGo Application")
	
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "server":
		startAPIServer()
	case "init":
		// Read in Database Config file
		config, err := config.ReadConfig("configs/dbconf.yml")
		if err != nil {
			log.Fatal(err)
		}

		// Start Database Initialization
		if err := db.InitDatabase(config); err != nil {
			log.Fatalf("Init failed: %v", err)
		}
	case "version":
		fmt.Fprintln(os.Stdout, "NPC Generator - v0.0a")
	case "help":
		printHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}
