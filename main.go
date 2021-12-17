package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sohhamm/todo-go-server/handlers"
)

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

func main() {
	r := initRouter()
	http.ListenAndServe(":8080", r)
}
