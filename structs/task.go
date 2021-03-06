package structs

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	// Task_id    int       `gorm:"primary_key"`
	gorm.Model

	Title     string    
	Time       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
	ProjectID  int
	StatusID    int  `gorm:"default:1"`
	Label   []Label `gorm:"many2many:task_labels;"`
}

// type TaskProject struct {
// 	TaskID    int `gorm:"primaryKey"`
// 	ProjectID int `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	DeletedAt gorm.DeletedAt
// }
type TaskLabel struct {
	TaskID    int `gorm:"primaryKey"`
	LabelID   int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

