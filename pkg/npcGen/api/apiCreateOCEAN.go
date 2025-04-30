package npcapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func APICreateOCEAN(context *gin.Context) {
	// Create new CS
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateOCEANData(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC OCEAN generation failed: %s", err)
		status := http.StatusInternalServerError
		context.JSON(status, Response{
			Status:    http.StatusText(status),
			Message:   msg,
			Timestamp: time.Now(),
		})
	}

	context.JSON(http.StatusOK, Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC OCEAN generated successfully",
		Data:      new_npc.OCEANToJSON(),
		Timestamp: time.Now(),
	})
}
