package models

import (
	"gorm.io/gorm"
)

type TodoModel struct {
	gorm.Model
	ID          int    `json:"id" gorm:"column:id;autoIncrement"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
