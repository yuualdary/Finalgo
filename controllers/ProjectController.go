package controllers

import (
	"FinalGo/config"
	"FinalGo/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProject(c *gin.Context) {


var(
	id = c.Params.ByName("id")
	User structs.User
	Project structs.Project
)
	if err := config.DB.Where("id = ?", id).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&Project)

	if err := config.DB.Omit("User").Create(&Project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, Project)
}

func GetAllProject(c *gin.Context) {

	var Project []structs.Project

	if err := config.DB.Find(&Project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, Project)

}

func UpdateProject(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		Project structs.Project
	)

	if err := config.DB.Find(&Project, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&Project)

	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
		return
	}

	c.JSON(http.StatusOK, Project)

}
