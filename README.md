# Docs & instructions

# 🧠 Go Task Manager API

A simple and clean RESTful Task Manager built with Go — perfect for learning Golang fundamentals by building a real project.

---

## 🚀 Features
- 📋 Create, complete, and delete tasks
- 💾 Store tasks in a local JSON file
- 🧪 Unit-tested core logic
- 🧱 Modular project structure
- 📡 Simple REST API (built with `net/http`)
- 🧑‍💻 Built from scratch without frameworks

---

## 📁 Project Structure

```text
go-task-manager/ 
├── api.go               # Main entry (starts the server)
├── go.mod               # Go module file
├── tasks.json           # Local file DB
│
├── task/                # Core domain logic
│   ├── model.go
│   ├── storage.go
│   └── task_test.go
│
├── handlers/            # HTTP route handlers
│   └── tasks.go
│
├── utils/               # Helper functions
│   └── id.go
│
├── auth/               # New! Handles JWT, login/signup
│   ├── auth.go
│   └── middleware.go

```
## 📡 API Endpoints (with cURL)

### Authentication
#### Signup
```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"username": "alawa", "password": "admin"}'
```
#### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "alawa", "password": "admin"}'
```
### 🆕 Create Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"title": "Buy groceries"}'
```
### 📋  Get All Tasks
```bash
curl -H "Authorization: Bearer <TOKEN>" http://localhost:8080/tasks
```
### ✅ Mark Task as Complete
```bash
curl -X PUT http://localhost:8080/tasks/1
```
### ❌ Delete Task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```
