package structs
import "gorm.io/gorm"

type Project struct {
	// Project_id   int `gorm:"primary_key"`
	gorm.Model
	Project_name string
	UserID int
}
