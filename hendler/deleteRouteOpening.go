package hendler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRouteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "Delete routeOpening",
	})
}
