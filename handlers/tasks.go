package handlers

import (
	"encoding/json"
	"go-task-manager/task"
	"net/http"
	"strconv"
	"strings"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task.Tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask task.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil || strings.TrimSpace(newTask.Title) == "" {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	newTask.ID = task.NextID
	task.NextID++
	task.Tasks = append(task.Tasks, newTask)
	task.SaveTasksToFile(task.Tasks)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for i := range task.Tasks {
		if task.Tasks[i].ID == id {
			task.Tasks[i].Completed = true
			task.SaveTasksToFile(task.Tasks)
			json.NewEncoder(w).Encode(task.Tasks[i])
			return
		}
	}
	http.NotFound(w, r)
	return
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for i := range task.Tasks {
		if task.Tasks[i].ID == id {
			task.Tasks = append(task.Tasks[:i], task.Tasks[i+1:]...)
			task.SaveTasksToFile(task.Tasks)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
	return
}
