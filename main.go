package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	c "github.com/sohhamm/todo-app-go-server/controllers"
	"github.com/sohhamm/todo-app-go-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func initRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	// add all routes
	s.HandleFunc("/todos", c.GetAllTodos).Methods("GET")
	s.HandleFunc("/todo/{id}", c.GetTodoByID).Methods("GET")
	s.HandleFunc("/todos", c.AddTodo).Methods("POST")
	s.HandleFunc("/todo/{id}", c.UpdateTodo).Methods("PUT")
	s.HandleFunc("/todo/{id}", c.DeleteTodo).Methods("DELETE")
	return r
}

func connectDatabase() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("host=%s sslmode=disable port=%s user=%s dbname=%s password=%s TimeZone=Asia/Kolkata", host, port, user, dbName, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to", dbName, "database")

	db.AutoMigrate(&models.TodoModel{})

}

func main() {
	r := initRouter()
	connectDatabase()

	http.ListenAndServe(":8080", r)
}
