package npcapi_reroll_npc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcapi "github.com/wholesomeow/npcGo/internal/api"
	namegen "github.com/wholesomeow/npcGo/internal/nameGen"
)

func APICreateName(context *gin.Context) {
	// Create new name
	new_name, err := namegen.CreateName(1)
	if err != nil {
		msg := fmt.Sprintf("NPC name generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}

	context.JSON(http.StatusOK, npcapi.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: "NPC name generated successfully",
		Data: map[string]string{
			"npc_name": new_name,
		},
		Timestamp: time.Now(),
	})
}
