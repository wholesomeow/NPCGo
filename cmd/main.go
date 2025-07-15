package main

import (
	"github.com/gin-gonic/gin"
	npcapi "github.com/wholesomeow/npcGo/internal/api"
	npcapi_npc "github.com/wholesomeow/npcGo/internal/api/npc"
	npcapi_reroll_enum "github.com/wholesomeow/npcGo/internal/api/reroll/enum"
	npcapi_reroll_npc "github.com/wholesomeow/npcGo/internal/api/reroll/npc"
)

func main() {
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
