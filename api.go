package main

import (
	"fmt"
	"go-task-manager/handlers"
	"go-task-manager/task"
	"go-task-manager/utils"
	"net/http"
)

func main() {
	task.Tasks = task.LoadTasksFromFile()
	task.NextID = utils.GetNextID(task.Tasks)

	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/", handlers.TaskByIDHandler)

	fmt.Println(" Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
