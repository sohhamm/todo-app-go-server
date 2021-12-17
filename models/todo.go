package models

type TodoModel struct {
	ID          int    `json:"id" gorm"column:id;autoIncrement"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
