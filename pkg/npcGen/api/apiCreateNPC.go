package npcapi

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	config "github.com/wholesomeow/npcGo/configs"
	npcgen "github.com/wholesomeow/npcGo/pkg/npcGen"
)

func apiCreateNPC(context *gin.Context) {
	// Read in Database Config file
	config, err := config.ReadConfig("configs/dbconf.yml")
	if err != nil {
		log.Fatal(err)
	}

	start_proc := time.Now()
	npc_object, err := npcgen.CreateNPC(config)
	if err != nil {
		log.Fatal(err)
	}
	end_proc := time.Now()
	elapsed_proc := end_proc.Sub(start_proc)
	log.Printf("npc created... elapsed time: %s", time.Duration.String(elapsed_proc))

	context.IndentedJSON(http.StatusOK, npc_object.DataToJSON())
}
