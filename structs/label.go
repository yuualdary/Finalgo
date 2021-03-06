package structs
import "gorm.io/gorm"
type Label struct {
	// Label_id  int `gorm:"primary_key"`
	gorm.Model
	LabelName string
	UserID     int

}
 

