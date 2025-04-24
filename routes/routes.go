package routes

import (
	"go-fruit-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/fruits/report-sugar", controllers.ReportSugar)
		api.GET("/fruits/load", controllers.LoadFruits)
	}
}
