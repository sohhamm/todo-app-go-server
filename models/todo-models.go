package models

import (
	"gorm.io/gorm"
)

/*-------- API Models --------*/

type TodoModel struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"desc"`
}
