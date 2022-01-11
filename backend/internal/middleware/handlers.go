package middleware

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/kennethk-1201/cvwo/backend/internal/helper"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

// Task Struct (Model)
type Task struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Body string `json:"body"`
	Deadline string `json:"deadline"`
}

// JsonResponse (Model)
type JsonResponse struct {
    Type    string `json:"type"`
    DataList    []Task `json:"datalist"`
    Data    Task `json:"data"`
    Message string `json:"message"`
}

// get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	db := helper.SetupDB()
	defer db.Close()

	fmt.Println("Getting tasks...")

	rows, err := db.Query("SELECT * FROM tasks")

	helper.CheckErr(err)

	var tasks []Task

	for rows.Next() {
        var id int
        var title string
        var body string
        var deadline string

        err = rows.Scan(&id, &title, &body, &deadline)

        // Check errors
        helper.CheckErr(err)

        tasks = append(tasks, Task{ID: id, Title: title, Body: body, Deadline: deadline})
    }

    var response = JsonResponse{Type: "success", DataList: tasks}

    json.NewEncoder(w).Encode(response)
}

// get specific task
func GetTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Methods", "GET")

	db := helper.SetupDB()
	defer db.Close()

	params := mux.Vars(r)

	fmt.Println("Getting task...")

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	row := db.QueryRow(`SELECT * FROM tasks where id=$1`, id)

	var taskId int
	var title string
	var body string
	var deadline string

	err = row.Scan(&taskId, &title, &body, &deadline)

	// Check errors
	helper.CheckErr(err)

    var response = JsonResponse{Type: "success", Data: Task{ID: taskId, Title: title, Body: body, Deadline: deadline}}

    json.NewEncoder(w).Encode(response)
}

// create task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	db := helper.SetupDB()
	defer db.Close()
}

// update task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	db := helper.SetupDB()
	defer db.Close()
}

// delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	db := helper.SetupDB()
	defer db.Close()
}
