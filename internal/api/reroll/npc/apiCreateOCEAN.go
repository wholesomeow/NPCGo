package npcapi_reroll_npc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcapi "github.com/wholesomeow/npcGo/internal/api"
	npcgen "github.com/wholesomeow/npcGo/internal/npcGen"
)

func APICreateOCEAN(context *gin.Context) {
	// Create empty NPC object and populate with CS Coordinates
	new_npc := npcgen.NPCBase{}
	new_npc.CS.Coords, _ = GetCSCoordinates(context)

	// Create new OCEAN Data
	err := npcgen.CreateOCEANData(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC OCEAN generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}

	// Return new data
	context.JSON(http.StatusOK, npcapi.Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC OCEAN generated successfully",
		Data:      new_npc.OCEANToJSON(),
		Timestamp: time.Now(),
	})
}
