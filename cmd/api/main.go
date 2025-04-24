package main

import (
	"go-fruit-api/database"
	"go-fruit-api/routes"
	"go-fruit-api/scheduler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	router := gin.Default()
	routes.RegisterRoutes(router)

	scheduler.StartScheduler()

	router.Run(":8080")
}
