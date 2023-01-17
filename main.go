package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("[*] Server listening on port ...")
	http.ListenAndServe(":8080", handlers.CORS()(router))
}