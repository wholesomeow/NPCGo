package npcapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	namegen "github.com/wholesomeow/npcGo/pkg/nameGen"
)

func APICreateName(context *gin.Context) {
	// Create new name
	new_name, err := namegen.CreateName()
	if err != nil {
		msg := fmt.Sprintf("NPC name generation failed: %s", err)
		status := http.StatusInternalServerError
		context.JSON(status, Response{
			Status:    http.StatusText(status),
			Message:   msg,
			Timestamp: time.Now(),
		})
	}

	context.JSON(http.StatusOK, Response{
		Status:  http.StatusText(http.StatusOK),
		Message: "NPC name generated successfully",
		Data: map[string]string{
			"npc_name": new_name,
		},
		Timestamp: time.Now(),
	})
}
