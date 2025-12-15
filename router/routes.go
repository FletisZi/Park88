package router

import (
	"github.com/fletiszi/goteste/hendler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	router.GET("/status", hendler.DatabaseStatus)
	v1.GET("/routeOpening", hendler.ShowRouteOpening)
	v1.POST("/routeOpening", hendler.CreateRouteOpening)
	v1.DELETE("/routeOpening", hendler.DeleteRouteOpening)
	v1.PUT("/routeOpening", hendler.UpdateRouteOpening)
	v1.GET("/routeOpenings", hendler.ListRouteOpening)
}
