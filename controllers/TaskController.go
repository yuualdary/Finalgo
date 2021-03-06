package controllers

import (
	"FinalGo/config"
	"FinalGo/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllTask godoc
// @Summary Show a Specific Task
// @Tags Task
// @Accept  json
// @Produce  json
// @success 200 {object} struct.Student
// @Router /task/{id} [get]
func GetTaskById(c *gin.Context) {
	var (
		id   = c.Params.ByName("id")
		Task []structs.Task
		// Project []structs.Project
	)
	// if err := config.DB.Joins("JOIN projects ON projects.project_id = tasks.project_id").Find(&Task, id).Error; err != nil {

	// if err := config.DB.Preload("tasks").First(&Project, id).Error; err != nil {
	// if err := config.DB.Preload("Projects").Preload("Projects").Preload("Projects").First(&Task, id).Error; err != nil {
	if err := config.DB.First(&Task, id).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, Task)

}

func GetAllTask(c *gin.Context) {

	var Task []structs.Task

	if err := config.DB.Find(&Task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, Task)

}

func CreateTask(c *gin.Context) {

	id := c.Params.ByName("id")
	var User structs.User
	var Task structs.Task

	if err := config.DB.Where("id = ?", id).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&Task)

	if err := config.DB.Omit("User", "Status", "Label").Create(&Task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, Task)
}
func UpdateTask(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		Task structs.Task
	)

	if err := config.DB.Find(&Task, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&Task)

	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&Task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
		return
	}

	c.JSON(http.StatusOK, Task)

}
func DeleteTask(c *gin.Context) {
	id := c.Params.ByName("id")

	var task structs.Task

	if err := config.DB.Where("id = ?", id).Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
}
