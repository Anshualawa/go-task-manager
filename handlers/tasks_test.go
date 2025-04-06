package handlers

import (
	"bytes"
	"go-task-manager/task"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Set up chi router for testing
func setupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/tasks", GetTasksHandler)
	r.Post("/tasks", CreateTaskHandler)
	r.Put("/tasks/{id}", UpdateTaskHandler)
	r.Delete("/tasks/{id}", DeleteTaskHandler)

	return r
}

func TestCreateAndGetTasks(t *testing.T) {
	// Reset tasks for test isolation
	task.Tasks = []task.Task{}
	task.NextID = 1

	router := setupRouter()

	// 1. Create a task
	body := []byte(`{"title":"Test Task via Chi"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", res.Code)
	}

	// 2. Get all tasks
	req = httptest.NewRequest(http.MethodGet, "/tasks", nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", res.Code)
	}

}
