package npcapi

import (
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
		log.Fatal(err)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("npc created... elapsed time: %s", time.Duration.String(elapsed_proc))

	context.JSON(http.StatusOK, npc_object.DataToJSON())
}
