package npcapi

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	namegen "github.com/wholesomeow/npcGo/pkg/nameGen"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func APICreateName(context *gin.Context) {
	start_proc := time.Now()
	var (
		err  error
		uuid = context.Param("uuid")
	)

	// Populate a new NPC object by querying database for UUID
	// and mapping returned data to new NPC object
	new_npc, err := npcgen.GetExistingNPC(uuid)
	if err != nil {
		log.Fatal(err)
	}

	// Create new name
	new_npc.Name, err = namegen.CreateName()
	if err != nil {
		log.Fatal(err)
	}

	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("npc created... elapsed time: %s", time.Duration.String(elapsed_proc))

	context.JSON(http.StatusOK, new_npc.NameToJSON())
}
