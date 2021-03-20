package apps

import (
	"github.com/one-piece-team1/one-piece-recovery/src/controllers/healths"
)

func MapUrls() {
	router.GET("/", healths.HealthCheck)
}
