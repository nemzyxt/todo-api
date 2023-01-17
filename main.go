/*
	Author 	  : Nemuel Wainaina
	Created   : 17th Jan 2023
	Modified  : 17th Jan 2023
*/

/*
	POST /tasks 	     : create a new task
	GET /tasks  	 	 : fetch all tasks
	GET /tasks/taskId 	 : fetch a specific task
	PUT /tasks/taskId 	 : update a specific task
	DELETE /tasks/taskId : delete a specific task
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func createTask(w http.ResponseWriter, r *http.Request) {

}

func getTasks(w http.ResponseWriter, r *http.Request) {
	
}

func getTask(w http.ResponseWriter, r *http.Request) {
	
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{taskId}", getTask).Methods("GET")
	router.HandleFunc("/tasks", updateTask).Methods("PUT")
	router.HandleFunc("/tasks", deleteTask).Methods("DELETE")

	fmt.Println("[*] Server listening on port ...")
	http.ListenAndServe(":8080", handlers.CORS()(router))
}