package hendler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRouteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "List All routeOpenings",
	})
}
