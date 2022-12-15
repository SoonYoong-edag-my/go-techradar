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

	router.GET("/techs", func(c *gin.Context) {
		log.Println("inside db var: ", db)
		var techs []model.Tech
		db.Order("category").Find(&techs)
		c.IndentedJSON(http.StatusOK, techs)
	})

	router.POST("/techs", func(c *gin.Context) {
		var newTech model.Tech
		if err := c.ShouldBindJSON(&newTech); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := newTech.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newTech)
		c.IndentedJSON(http.StatusCreated, newTech)
	})

	router.PUT("/techs/:id", func(c *gin.Context) {
		var existingTech, updateTech model.Tech

		if err := db.First(&existingTech, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := c.ShouldBindJSON(&updateTech); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := updateTech.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateTech.Id = existingTech.Id
		// have to use below instead of db.Model(&existingTech).Updates(updateTech) as it doesn't update to false or null value
		db.Model(&existingTech).Updates(map[string]interface{}{"Category": updateTech.Category, "Status": updateTech.Status,
			"Name": updateTech.Name, "Active": updateTech.Active, "Moved": updateTech.Moved})
		c.IndentedJSON(http.StatusOK, existingTech)
	})

	router.DELETE("/techs/:id", func(c *gin.Context) {
		var existingTodo model.Tech
		if err := db.First(&existingTodo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		db.Model(&existingTodo).Delete(&model.Tech{}, existingTodo.Id)
	})
}
