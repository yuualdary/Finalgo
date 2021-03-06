package controllers

import (
	"FinalGo/config"
	"FinalGo/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLabel(c *gin.Context) {

	var (
		id = c.Params.ByName("id")
	 	User  structs.User
	 	Label structs.Label	
	)
	// c.ShouldBindJSON(&Label)

	// if err := config.DB.Create(&Label).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error Item": "Record can't create!",
	// 	})
	// }

	// c.JSON(http.StatusOK, Label)


	// id := c.Params.ByName("id")
	// var user models.User
	// var label models.Label

	if err := config.DB.Where("id = ?", id).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&Label)

	if err := config.DB.Omit("User").Create(&Label).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, Label)
}

func GetAllLabel(c *gin.Context) {

	var Label []structs.Label

	if err := config.DB.Find(&Label).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, Label)

}

func UpdateLabel(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		Label structs.Label
	)

	if err := config.DB.Find(&Label, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&Label)

	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Label).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
		return
	}

	c.JSON(http.StatusOK, Label)

}
