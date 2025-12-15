package hendler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowRouteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET routeOpening",
	})
}
