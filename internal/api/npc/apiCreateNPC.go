package npcapi_npc

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
	npcapi "github.com/wholesomeow/npcGo/pkg/npcGen/api"
)

func APICreateNPC(context *gin.Context) {
	start_proc := time.Now()
	npc_object, err := npcgen.CreateNPC(1)
	if err != nil {
		msg := fmt.Sprintf("NPC name generation failed: %s", err)
		status, response := npcapi.Response500(msg)
		context.JSON(status, response)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("npc created... elapsed time: %s", time.Duration.String(elapsed_proc))

	context.JSON(http.StatusOK, npcapi.Response{
		Status:    http.StatusText(http.StatusOK),
		Message:   "NPC name generated successfully",
		Data:      npc_object.DataToJSON(),
		Timestamp: time.Now(),
	})
}
