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
	npc_group.GET("/createNPC", npcapi.APICreateNPC)

	router.Run("0.0.0.0:8080")
}
