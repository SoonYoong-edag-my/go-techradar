package controllers

import (
	"github.com/gin-gonic/gin"
	"go-techradar/model"
	"log"
	"net/http"
)

func Init(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/blips", func(c *gin.Context) {
		log.Println("inside db var: ", db)
		var blips []model.Blip
		db.Order("category").Find(&blips)
		c.IndentedJSON(http.StatusOK, blips)
	})

	router.POST("/blips", func(c *gin.Context) {
		var newBlip model.Blip
		if err := c.ShouldBindJSON(&newBlip); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := newBlip.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newBlip)
		c.IndentedJSON(http.StatusCreated, newBlip)
	})

	router.PUT("/blips/:id", func(c *gin.Context) {
		var existingBlip, updateBlip model.Blip

		if err := db.First(&existingBlip, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := c.ShouldBindJSON(&updateBlip); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := updateBlip.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateBlip.Id = existingBlip.Id
		// have to use below instead of db.Model(&existingBlip).Updates(updateBlip) as it doesn't update to false or null value
		db.Model(&existingBlip).Updates(map[string]interface{}{"Category": updateBlip.Category, "Status": updateBlip.Status,
			"Name": updateBlip.Name, "Active": updateBlip.Active, "Moved": updateBlip.Moved, "Description": updateBlip.Description})
		c.IndentedJSON(http.StatusOK, existingBlip)
	})

	router.DELETE("/blips/:id", func(c *gin.Context) {
		var existingTodo model.Blip
		if err := db.First(&existingTodo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		db.Model(&existingTodo).Delete(&model.Blip{}, existingTodo.Id)
	})
}
