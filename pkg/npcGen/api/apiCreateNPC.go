package npcapi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func APICreateNPC(context *gin.Context) {
	start_proc := time.Now()
	npc_object, err := npcgen.CreateNPC()
	if err != nil {
		msg := fmt.Sprintf("NPC name generation failed: %s", err)
		status := http.StatusInternalServerError
		context.JSON(status, Response{
			Status:    http.StatusText(status),
			Message:   msg,
			Timestamp: time.Now(),
		})
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("npc created... elapsed time: %s", time.Duration.String(elapsed_proc))

	context.JSON(http.StatusOK, Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC name generated successfully",
		Data:      npc_object.DataToJSON(),
		Timestamp: time.Now(),
	})
}
