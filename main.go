package main

import (
	"FinalGo/config"
	"FinalGo/controllers"

	// _ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// _ "FinalGo"
	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)

// @title Todo API
// @version 1.0
// @description This is a basic API Todo using Gin and Gorm.

// @contact.name API Support
// @contact.email wahyuyudistiro@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {

	router := gin.Default()

	config.ConnectDatabase()

	//route
	v1 := router.Group("/api/v1")
	{
		v1.GET("/task/:id", controllers.GetTaskById)
		v1.POST("/createtask/:id", controllers.CreateTask)
		v1.PUT("update/:id", controllers.UpdateTask)
		v1.GET("/getalltask", controllers.GetAllTask)
		v1.DELETE("/deletetask/:id", controllers.DeleteTask)

		v1.POST("/createproject/:id", controllers.CreateProject)
		v1.PUT("/updateproject/:id", controllers.UpdateProject)
		v1.GET("/project", controllers.GetAllProject)
		// v1.DELETE("/deleteproject/:id", controllers.De)

		v1.POST("/createlabel/:id", controllers.CreateLabel)
		v1.PUT("/updatelabel/:id", controllers.UpdateLabel)
		v1.GET("/label", controllers.GetAllLabel)

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
