package hendler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRouteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "Update routeOpening",
	})
}
