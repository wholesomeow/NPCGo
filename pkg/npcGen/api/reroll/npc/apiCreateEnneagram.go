package npcapi_reroll_npc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
	npcapi "github.com/wholesomeow/npcGo/pkg/npcGen/api"
)

func APICreateEnneagram(context *gin.Context) {
	// Create new Enneagram
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateEnneagram(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC Enneagram generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}

	context.JSON(http.StatusOK, npcapi.Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC Enneagram generated successfully",
		Data:      new_npc.EnneagramToJSON(),
		Timestamp: time.Now(),
	})
}
