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
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // local default
	}

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


	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:3000"}, // or "*" for dev
		AllowedOrigins:   []string{"http://localhost:3000", "https://your-app-url.up.railway.app"},
		AllowedMethods:   []string{"GET", POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true
	}))

	r.Use(auth.LoggingMiddleware)

	r.Post("/signup", auth.SignupHandler)
	r.Post("/login", auth.LoginHandler)


	// 🔐 Protect task Routes
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware) // 🔐 apply to all below

		r.Get("/tasks", handlers.GetTasksHandler)
		r.Post("/tasks", handlers.CreateTaskHandler)
		r.Put("/tasks/{id}", handlers.UpdateTaskHandler)
		r.Delete("/tasks/{id}", handlers.DeleteTaskHandler)

	})

	fmt.Println(" Server running at http://localhost:" + port)
	fmt.Println(" Server running on port " + port)
	http.ListenAndServe(":"+port, r)
}
