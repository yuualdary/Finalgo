package structs

import "gorm.io/gorm"
type Status struct {
	// Status_id  int `gorm:"primary_key"`
	gorm.Model
	StatusName string
}
