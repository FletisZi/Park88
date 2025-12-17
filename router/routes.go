package router

import (
	"github.com/fletiszi/goteste/hendler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	router.LoadHTMLGlob("templates/**/*")

	router.GET("/", hendler.Home)
	router.GET("/new-estacionamento", hendler.PageCreateEstacionamentos)

	v1 := router.Group("/api/v1")

	router.GET("/status", hendler.DatabaseStatus)
	router.GET("/dbstatus", hendler.DBStats)
	v1.GET("/routeOpening", hendler.ShowRouteOpening)
	v1.POST("/routeOpening", hendler.CreateRouteOpening)
	v1.DELETE("/routeOpening", hendler.DeleteRouteOpening)
	v1.PUT("/routeOpening", hendler.UpdateRouteOpening)
	v1.GET("/routeOpenings", hendler.ListRouteOpening)
	v1.GET("/estacionamentos", hendler.ListEstacionamentos)
	v1.POST("/estacionamentos", hendler.CreateEstacionamentos)
	v1.DELETE("/estacionamentos/:id", hendler.DeleteEstacionamentos)
	v1.PUT("estacionamentos/:id", hendler.UpdateEstacionamentos)
}
