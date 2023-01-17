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

func main() {
	router := mux.NewRouter()

	fmt.Println("[*] Server listening on port ...")
	http.ListenAndServe(":8080", handlers.CORS()(router))
}