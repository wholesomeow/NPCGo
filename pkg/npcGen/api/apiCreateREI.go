package npcapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func APICreateREI(context *gin.Context) {
	// Create new RE
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateREIData(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC REI generation failed: %s", err)
		status := http.StatusInternalServerError
		context.JSON(status, Response{
			Status:    http.StatusText(status),
			Message:   msg,
			Timestamp: time.Now(),
		})
	}

	context.JSON(http.StatusOK, Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC REI generated successfully",
		Data:      new_npc.REIToJSON(),
		Timestamp: time.Now(),
	})
}
