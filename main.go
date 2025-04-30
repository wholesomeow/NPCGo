package main

import (
	"github.com/gin-gonic/gin"
	npcapi "github.com/wholesomeow/npcGo/pkg/npcGen/api"
)

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	npc_group := router.Group("/npc")
	npc_group.GET("/createNPC/", npcapi.APICreateNPC)
	npc_group.GET("/createName/", npcapi.APICreateName)
	npc_group.GET("/createCS/", npcapi.APICreateCS)
	npc_group.GET("/createOCEAN/", npcapi.APICreateOCEAN)
	npc_group.GET("/createMICE/", npcapi.APICreateMICE)
	npc_group.GET("/createREI/", npcapi.APICreateREI)
	npc_group.GET("/createEnneagram/", npcapi.APICreateEnneagram)

	router.Run("0.0.0.0:8080")
}
