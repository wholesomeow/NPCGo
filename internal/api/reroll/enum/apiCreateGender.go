package npcapi_reroll_enum

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcapi "github.com/wholesomeow/npcGo/internal/api"
	npcgen "github.com/wholesomeow/npcGo/internal/npcGen"
)

func APICreateGender(context *gin.Context) {
	// Create new CS
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateGenderType(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC Gender generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}

	context.JSON(http.StatusOK, npcapi.Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC Gender generated successfully",
		Data:      new_npc.GenderToJSON(),
		Timestamp: time.Now(),
	})
}
