package web

import (
	"dc-monitor/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.gohtml")
	router.GET("/", handler.DashboardHandler)
	router.GET("/api/all", handler.APIDashboardHandler)
	router.GET("/api/gh/:id", handler.APIGHHandler)
	router.Run()

}
