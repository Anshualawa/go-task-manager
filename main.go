package main

import (
	"fmt"
	"go-task-manager/auth"
	"go-task-manager/handlers"
	"go-task-manager/task"
	"go-task-manager/utils"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	task.Tasks = task.LoadTasksFromFile()
	task.NextID = utils.GetNextID(task.Tasks)

	r := chi.NewRouter()

	// ✅ Add CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*","https://reactjs-task-manager-production.up.railway.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// 🛡️ Your middlewares
	//r.Use(auth.LoggingMiddleware)

	// 🔐 Auth routes
	r.Post("/signup", auth.SignupHandler)
	r.Post("/login", auth.LoginHandler)

	// ✅ Task routes (protected)
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Get("/tasks", handlers.GetTasksHandler)
		r.Post("/tasks", handlers.CreateTaskHandler)
		r.Put("/tasks/{id}", handlers.UpdateTaskHandler)
		r.Delete("/tasks/{id}", handlers.DeleteTaskHandler)
	})

	fmt.Println("🚀 Server running at http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}
