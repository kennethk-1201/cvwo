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

// get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
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

    response := struct {
		Type string `json:"type"`
		Data []Task `json:"data"`
		Message string `json:"message"`
	}{
		Type: "success",
		Data: tasks,
		Message: "Returned tasks!",
	}

    json.NewEncoder(w).Encode(response)
}

// get specific task
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")

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

    response := struct {
		Type string `json:"type"`
		Data Task `json:"data"`
		Message string `json:"message"`
	}{
		Type: "success",
		Data: Task{ID: taskId, Title: title, Body: body, Deadline: deadline},
		Message: "Returned task!",
	}

    json.NewEncoder(w).Encode(response)
}

// create task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	var task Task

	err1 := json.NewDecoder(r.Body).Decode(&task)
	helper.CheckErr(err1)

	db := helper.SetupDB()
	defer db.Close()

	sqlStatement := `INSERT INTO tasks (title, body, deadline) VALUES ($1, $2, $3) RETURNING id`

	// the inserted id will store in this id
    var id int64

    // execute the sql statement
    // Scan function will save the insert id in the id
    err2 := db.QueryRow(sqlStatement, task.Title, task.Body, task.Deadline).Scan(&task.ID)
	helper.CheckErr(err2)

	fmt.Printf("Created new task with id %v", id)

	response := struct {
		Type string `json:"type"`
		Data Task `json:"data"`
		Message string `json:"message"`
	}{
		Type: "success",
		Data: task,
		Message: "Created a task!",
	}

    json.NewEncoder(w).Encode(response)
}

func (t *Task) modifyTaskID(i int) {
	t.ID = i
}

// update task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task Task
	
	err1 := json.NewDecoder(r.Body).Decode(&task)
	helper.CheckErr(err1)

	db := helper.SetupDB()
	defer db.Close()

	params := mux.Vars(r)

	fmt.Println("Getting task...")

	sqlStatement := `UPDATE tasks SET title=$2, body=$3, deadline=$4 WHERE id=$1`

	// convert the id type from string to int
	id, err2 := strconv.Atoi(params["id"])
	helper.CheckErr(err2)

	_, err3 := db.Exec(sqlStatement, id, task.Title, task.Body, task.Deadline)
	helper.CheckErr(err3)

	fmt.Printf("Updated task with id %v", id)

	task.modifyTaskID(id)

	response := struct {
		Type string `json:"type"`
		Data Task `json:"data"`
		Message string `json:"message"`
	}{
		Type: "success",
		Data: task,
		Message: "Updated a task!",
	}

    json.NewEncoder(w).Encode(response)
}

// delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	db := helper.SetupDB()
	defer db.Close()

	// get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to int
    id, err1 := strconv.Atoi(params["id"])

    helper.CheckErr(err1)

	sqlStatement := `DELETE FROM tasks WHERE id=$1`

	// execute the sql statement
    _, err2 := db.Exec(sqlStatement, id)

	helper.CheckErr(err2)

	fmt.Printf("Deleted task with id %v", id)

	// format the reponse message
    response := struct {
        Type string `json:"type"`
        ID int `json:"id"`
		Message string `json:"message"`
    }{
		Type: "success",
		ID: id,
		Message: "Deleted a task!",
	}

    // send the response
    json.NewEncoder(w).Encode(response)
}
