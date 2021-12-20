package models

import (
	"gorm.io/gorm"
)

/*-------- API Models --------*/

type Todo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"desc"`
}
