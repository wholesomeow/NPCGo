package npcapi_reroll_enum

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
	npcapi "github.com/wholesomeow/npcGo/pkg/npcGen/api"
)

func APICreateType(context *gin.Context) {
	// Create new CS
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateNPCType(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC Type generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}

	context.JSON(http.StatusOK, npcapi.Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC Type generated successfully",
		Data:      new_npc.TypeToJSON(),
		Timestamp: time.Now(),
	})
}
