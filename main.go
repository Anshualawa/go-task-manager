package main

import (
	"fmt"
	"go-task-manager/handlers"
	"go-task-manager/task"
	"go-task-manager/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load existing tasks and setup ID
	task.Tasks = task.LoadTasksFromFile()
	task.NextID = utils.GetNextID(task.Tasks)

	// Init chi router
	r := chi.NewRouter()

	// 🔥 Built-in middleware
	r.Use(middleware.Logger)    // Logs requests
	r.Use(middleware.RequestID) // Add unique request ID to each req
	r.Use(middleware.RealIP)    // Logs real IP (useful in prod)
	r.Use(middleware.Recoverer) // Recovers from panics & logs stack trace

	// Routes
	r.Get("/tasks", handlers.GetTasksHandler)
	r.Post("/tasks", handlers.CreateTaskHandler)
	r.Put("/tasks/{id}", handlers.UpdateTaskHandler)
	r.Delete("/tasks/{id}", handlers.DeleteTaskHandler)

	fmt.Println(" Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
