package controllers

import (
	"go-fruit-api/database"
	"go-fruit-api/models"
	"go-fruit-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReportSugar(c *gin.Context) {
	var fruits []models.Fruit
	database.DB.Find(&fruits)

	highSugar := []gin.H{}
	lowSugar := []gin.H{}

	for _, f := range fruits {
		item := gin.H{"id": f.ID, "name": f.Name}
		if f.Sugar >= 10 {
			highSugar = append(highSugar, item)
		} else {
			lowSugar = append(lowSugar, item)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"high_sugar":       highSugar,
		"low_sugar":        lowSugar,
		"total_high_sugar": len(highSugar),
		"total_low_sugar":  len(lowSugar),
	})
}

func LoadFruits(c *gin.Context) {
	err := services.LoadFruitsFromAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fruits loaded successfuly"})
}
