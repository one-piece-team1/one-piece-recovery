package healths

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Health = "Recovery Server is Healthly"
)

func HealthCheck(ctx *gin.Context) {
	ctx.String(http.StatusOK, Health)
}
