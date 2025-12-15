package hendler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "Create routeOpening",
	})
}
