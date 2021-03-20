package apps

import (
	"github.com/gin-gonic/gin"
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

	if err := router.Run(":7076"); err != nil {
		panic(err)
	}
}
