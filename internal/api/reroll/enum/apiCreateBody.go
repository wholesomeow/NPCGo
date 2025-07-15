package npcapi_reroll_enum

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcapi "github.com/wholesomeow/npcGo/internal/api"
	npcgen "github.com/wholesomeow/npcGo/internal/npcGen"
)

func APICreateBody(context *gin.Context) {
	// Create new CS
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateBodyType(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC Body generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}

	context.JSON(http.StatusOK, npcapi.Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC Body generated successfully",
		Data:      new_npc.BodyToJSON(),
		Timestamp: time.Now(),
	})
}
