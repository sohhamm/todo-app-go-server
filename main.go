package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//handler
	fmt.Fprintf(w, "Hello this is a basic server")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
