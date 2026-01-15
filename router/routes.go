package router

import (
	"time"

	"github.com/fletiszi/goteste/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // depois pode restringir
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.LoadHTMLGlob("templates/**/*")

	router.GET("/", handler.Home)
	router.GET("/new-estacionamento", handler.PageCreateEstacionamentos)
	router.GET("/vagas-estacionamento", handler.PageVagasEstacionamento)
	router.GET("/home", handler.PageUpdateEstacionamentos)
	router.GET("/create-vaga", handler.PageCreateVaga)

	v1 := router.Group("/api/v1")

	router.GET("/status", handler.DatabaseStatus)
	router.GET("/dbstatus", handler.DBStats)
	v1.GET("/routeOpening", handler.ShowRouteOpening)
	v1.POST("/routeOpening", handler.CreateRouteOpening)
	v1.DELETE("/routeOpening", handler.DeleteRouteOpening)
	v1.PUT("/routeOpening", handler.UpdateRouteOpening)
	v1.GET("/routeOpenings", handler.ListRouteOpening)

	v1.GET("/estacionamentos", handler.ListEstacionamentos)
	v1.POST("/estacionamentos", handler.CreateEstacionamentos)
	v1.DELETE("/estacionamentos/:id", handler.DeleteEstacionamentos)
	v1.PUT("estacionamentos/:id", handler.UpdateEstacionamentos)
	v1.GET("/estacionamentos/:id/vagas", handler.GetVagasStatus)
	v1.POST("/vagas", handler.CreateVagas)

}
