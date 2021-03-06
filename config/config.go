package config

import (
	"FinalGo/structs"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var (
		Task = structs.Task{}
		// TaskProject = structs.TaskProject{}
		// User = structs.User{}
		UserTask = structs.UserTask{}
		Label     = structs.Label{}
		TaskLabel = structs.TaskLabel{}
	)
	dsn := "root:@tcp(127.0.0.1:3306)/final?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	// db.SetupJoinTable(Task, "Projects", TaskProject)

	db.SetupJoinTable(Task, "Users", UserTask)
	db.SetupJoinTable(Label, "Tasks", TaskLabel)


	db.AutoMigrate(&structs.Label{}, &structs.Project{}, &structs.Status{}, &structs.Task{},&structs.User{})

	fmt.Println("Connected to Database")

	DB = db
}
