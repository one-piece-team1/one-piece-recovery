package apps

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/one-piece-team1/one-piece-recovery/src/config"
	"github.com/one-piece-team1/one-piece-recovery/src/pkg/postgres"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	client := postgres.DBClient{}
	client.PgConnect()

	// url mapping but only health check here
	MapUrls()

	port := config.GetPort()
	appPort := fmt.Sprintf(":" + port)
	if err := router.Run(appPort); err != nil {
		panic(err)
	}
}
