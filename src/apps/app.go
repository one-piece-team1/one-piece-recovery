package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/one-piece-team1/one-piece-recovery/src/config"
	"log"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// url mapping but only health check here
	MapUrls()

	port := config.GetPort()
	appPort := fmt.Sprintf(":" + port)
	if err := router.Run(appPort); err != nil {
		panic(err)
	}
}
