package npcapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func APICreateMICE(context *gin.Context) {
	// Create new MICE
	new_npc := npcgen.NPCBase{}
	err := npcgen.CreateMICEData(&new_npc)
	if err != nil {
		msg := fmt.Sprintf("NPC MICE generation failed: %s", err)
		status := http.StatusInternalServerError
		context.JSON(status, Response{
			Status:    http.StatusText(status),
			Message:   msg,
			Timestamp: time.Now(),
		})
	}

	context.JSON(http.StatusOK, Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC MICE generated successfully",
		Data:      new_npc.MICEToJSON(),
		Timestamp: time.Now(),
	})
}
