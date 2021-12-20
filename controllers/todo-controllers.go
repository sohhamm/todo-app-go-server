package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/sohhamm/todo-app-go-server/models"
	"gorm.io/gorm"
)

/*-------- API Controllers --------*/

func GetAllTodos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todos []m.Todo
		dbTodos := db.Find(&todos)
		err := dbTodos.Error
		if err != nil {
			panic("failed to get  all todos")
		}
		json.NewEncoder(w).Encode(&todos)
	}

}

func GetTodoByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var todo m.Todo
		dbTodo := db.First(&todo, params["id"])
		err := dbTodo.Error
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(&todo)

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
		params := mux.Vars(r)
		var todo m.Todo
		json.NewDecoder(r.Body).Decode(&todo)

		dbTodo := db.First(&todo, params["id"])
		err := dbTodo.Error
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(&todo)
	}
}

func DeleteTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
