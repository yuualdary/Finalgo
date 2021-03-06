package structs

import 	"gorm.io/gorm"

type User struct{

	// User_id int `gorm:"primary_key"`
	gorm.Model
	User_name string
	Task []Task `gorm:"many2many:user_tasks;"`
	Project []Project
	Label []Label

}

type UserTask struct {
	UserID int `gorm:"primaryKey"`
	TaskID int `gorm:"primaryKey"`
}
