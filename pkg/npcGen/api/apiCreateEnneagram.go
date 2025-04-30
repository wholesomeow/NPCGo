package npcapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func APICreateEnneagram(context *gin.Context) {
	// Create new Enneagram
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateEnneagram(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC Enneagram generation failed: %s", err)
		status := http.StatusInternalServerError
		context.JSON(status, Response{
			Status:    http.StatusText(status),
			Message:   msg,
			Timestamp: time.Now(),
		})
	}

	context.JSON(http.StatusOK, Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC Enneagram generated successfully",
		Data:      new_npc.EnneagramToJSON(),
		Timestamp: time.Now(),
	})
}
