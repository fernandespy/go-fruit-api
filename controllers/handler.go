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
		item := gin.H{"fruit_id": f.ID, "name": f.Name}
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

func GetFruits(c *gin.Context) {
	var fruits []models.Fruit
	if err := database.DB.Find(&fruits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR getting fruits"})
		return
	}

	if len(fruits) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Fruits empty"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Fruits": fruits,
	})
}

func GetFruitsByID(c *gin.Context) {
	fruitID := c.Param("fruit_id")
	var fruit models.Fruit

	if err := database.DB.Unscoped().Where("fruit_id = ?", fruitID).First(&fruit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fruit not found"})
		return
	}

	if fruit.DeletedAt.Valid {
		c.JSON(http.StatusOK, gin.H{"Fruit was deleted_at": fruit.DeletedAt})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Fruit": fruit,
	})
}

func DeleteFruitByID(c *gin.Context) {
	fruitID := c.Param("fruit_id")
	var fruit models.Fruit

	if err := database.DB.Unscoped().Where("fruit_id = ?", fruitID).First(&fruit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fruit not found"})
		return
	}

	if fruit.DeletedAt.Valid {
		c.JSON(http.StatusOK, gin.H{"Fruit already deleted:": fruit.DeletedAt})
		return
	}

	if err := database.DB.Delete(&fruit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete fruit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fruit deleted successfully"})
}
