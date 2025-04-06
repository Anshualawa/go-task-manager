package task

// Task struct to hold task details
type Task struct {
	ID        int
	Title     string
	Completed bool
}

var Tasks []Task
var NextID int
