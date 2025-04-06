package task

import "testing"

func TestAddTask(t *testing.T) {
	var testTasks []Task
	newTask := Task{ID: 1, Title: "Learn Go"}
	testTasks = append(testTasks, newTask)

	if len(testTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(testTasks))
	}
}
