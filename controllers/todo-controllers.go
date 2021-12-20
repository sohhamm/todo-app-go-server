package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/sohhamm/todo-app-go-server/models"
	"gorm.io/gorm"
)

/*-------- API Controllers --------*/

func GetAllTodos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo m.Todo
		json.NewDecoder(r.Body).Decode(&todo)
		fmt.Println(todo)
		createdTodo := db.Create(&todo)
		err := createdTodo.Error
		if err != nil {
			panic("failed to add todo")
		}
		json.NewEncoder(w).Encode(&todo)
	}

}

func GetTodoByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func AddTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo m.Todo
		json.NewDecoder(r.Body).Decode(&todo)
		createdTodo := db.Create(&todo)
		err := createdTodo.Error
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(&todo)
	}
}

func UpdateTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
