package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sohhamm/todo-app-go-server/handlers"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func initRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	// add all routes
	s.HandleFunc("/todos", handlers.GetAllTodosHandler).Methods("GET")
	s.HandleFunc("/todo/{id}", handlers.GetTodoByIDHandler).Methods("GET")
	s.HandleFunc("/todos", handlers.AddTodoHandler).Methods("POST")
	s.HandleFunc("/todo/{id}", handlers.UpdateTodoHandler).Methods("PUT")
	s.HandleFunc("/todo/{id}", handlers.DeleteTodoHandler).Methods("DELETE")
	return r
}

func connectDatabase() {
	// dialect := os.Getenv("DIALECT")
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

	fmt.Println(db)

	fmt.Println("Connected to the database successfully")

}

func main() {
	r := initRouter()

	http.ListenAndServe(":8080", r)
}
