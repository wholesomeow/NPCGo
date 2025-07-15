package npcapi_reroll_npc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	namegen "github.com/wholesomeow/npcGo/pkg/nameGen"
	npcapi "github.com/wholesomeow/npcGo/pkg/npcGen/api"
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
