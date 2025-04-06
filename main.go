package main

import (
	"fmt"
	"go-task-manager/task"
)

// func main() {
// 	// Create a task
// 	// task1 := Task{
// 	// 	ID:        1,
// 	// 	Title:     "Learn Go structs",
// 	// 	Completed: false,
// 	// }

// 	// // Print task details
// 	// fmt.Println("Task :", task1)

// 	// // Create multiple tasks
// 	// tasks := []Task{
// 	// 	{ID: 1, Title: "Learn Go structs", Completed: false},
// 	// 	{ID: 2, Title: "Practice loops", Completed: false},
// 	// 	{ID: 3, Title: "Celebrate progress", Completed: false},
// 	// }

// 	// // Print all tasks
// 	// fmt.Println("📝 Task List :")

// 	// for _, task := range tasks {
// 	// 	status := "❌"
// 	// 	if task.Completed {
// 	// 		status = "✅"
// 	// 	}
// 	// 	fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
// 	// }

// 	// // Mark task with ID 2 as complete
// 	// tasks = MarkTaskComplete(tasks, 2)

// 	// var tasks []Task
// 	// nextID := 1

// 	// tasks, nextID = CreateTask(tasks, "Learn Go functions", nextID)
// 	// tasks, nextID = CreateTask(tasks, "Add delete feature", nextID)
// 	// tasks, nextID = CreateTask(tasks, "Practice hard", nextID)

// 	// tasks = MarkTaskComplete(tasks, 2)
// 	// tasks = DeleteTask(tasks, 1)

// 	// PrintTasks(tasks)

// 	// var tasks []Task
// 	// nextID := 1
// 	// reader := bufio.NewReader(os.Stdin)

// 	tasks := loadTasksFromFile()
// 	nextID := getNextID(tasks)
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Println("\nChoose an option :")
// 		fmt.Println("1. Add task")
// 		fmt.Println("2. Delete task")
// 		fmt.Println("3. Complete task")
// 		fmt.Println("4. Show task")
// 		fmt.Println("5. Exit")

// 		fmt.Print("👉 Enter choice : ")
// 		choiceInput, _ := reader.ReadString('\n')
// 		choice := strings.TrimSpace(choiceInput)

// 		switch choice {
// 		case "1":
// 			fmt.Print("📝Enter task title : ")
// 			title, _ := reader.ReadString('\n')
// 			title = strings.TrimSpace(title)
// 			tasks, nextID = CreateTask(tasks, title, nextID)
// 			saveTasksToFile(tasks)
// 		case "2":
// 			fmt.Print("🗑️Enter task ID to delete : ")
// 			idInput, _ := reader.ReadString('\n')
// 			id, _ := strconv.Atoi(strings.TrimSpace(idInput))
// 			tasks = DeleteTask(tasks, id)
// 			saveTasksToFile(tasks)
// 		case "3":
// 			fmt.Print("✅Enter task ID to mark complete : ")
// 			idInput, _ := reader.ReadString('\n')
// 			id, _ := strconv.Atoi(strings.TrimSpace(idInput))
// 			tasks = MarkTaskComplete(tasks, id)
// 			saveTasksToFile(tasks)
// 		case "4":
// 			PrintTasks(tasks)
// 		case "5":
// 			fmt.Println("👋Exiting. Bye!")
// 			return
// 		default:
// 			fmt.Println("❓Invalid option. Try again.")
// 		}
// 	}

// }

// MarkTaskComplete marks a task as complete by ID
func MarkTaskComplete(tasks []task.Task, id int) []task.Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Printf("🎉 Task '%s' marked as complete!\n", task.Title)
		}
	}
	return tasks
}

// CreateTask adds a new task to the list
func CreateTask(tasks []task.Task, title string, nextID int) ([]task.Task, int) {
	newTask := task.Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	fmt.Printf("➕ Task '%s' added with ID %d\n", title, nextID)
	return tasks, nextID + 1
}

// DeleteTask removes a task by ID
func DeleteTask(tasks []task.Task, id int) []task.Task {
	for i, task := range tasks {
		if task.ID == id {
			fmt.Printf("❌ Task '%s' deleted!\n", task.Title)
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	fmt.Printf("⚠️ Task with ID %d not found\n", id)
	return tasks
}

// PrintTasks displays the task list
func PrintTasks(tasks []task.Task) {
	fmt.Println("\n📃 Current Task List :")
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}
