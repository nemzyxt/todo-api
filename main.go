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
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"todo/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
type Task models.Task

func initDB() {
	db_name := os.Getenv("DB_NAME")
	db, _ = gorm.Open(sqlite.Open(db_name))
	db.AutoMigrate(Task{})
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	db.Create(&task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&task)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	if(len(tasks) > 0) {
		json.NewEncoder(w).Encode(&tasks)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, _ := strconv.Atoi(params["id"])
	var task Task
	db.Find(&task, Task{Id: taskId})
	if task != (Task{}) {
		json.NewEncoder(w).Encode(&task)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, _ := strconv.Atoi(params["id"])
	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	db.Where(Task{Id: taskId}).Updates(Task{Title: task.Title, Description: task.Description})	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, _ := strconv.Atoi(params["id"])
	db.Delete(Task{Id: taskId})
	w.Write([]byte("deleted"))
}

func main() {
	router := mux.NewRouter()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	initDB()
	fmt.Println("[*] Server listening on port ...")
	http.ListenAndServe(":8080", handlers.CORS()(router))
}