package npcapi_reroll_npc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcapi "github.com/wholesomeow/npcGo/internal/api"
	npcgen "github.com/wholesomeow/npcGo/internal/npcGen"
)

func APICreateREI(context *gin.Context) {
	// Create empty NPC object and populate with CS Coordinates
	new_npc := npcgen.NPCBase{}
	new_npc.CS.Coords, _ = GetCSCoordinates(context)

	// Create new REI
	err := npcgen.CreateREIData(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC REI generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}

	context.JSON(http.StatusOK, npcapi.Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC REI generated successfully",
		Data:      new_npc.REIToJSON(),
		Timestamp: time.Now(),
	})
}
