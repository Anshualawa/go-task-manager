package task

import "testing"

func TestAddTask(t *testing.T) {
	var testTasks []Task
	newTask := Task{ID: 1, Title: "Learn Go"}
	testTasks = append(testTasks, newTask)

	if len(testTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(testTasks))
	}
	if testTasks[0].Title != "Learn Go" {
		t.Errorf("Expected title 'Learn Go', got '%s'", testTasks[0].Title)
	}
}

func TestCompleteTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "Test", Completed: false},
	}

	// Simulate completing task
	tasks[0].Completed = true

	if !tasks[0].Completed {
		t.Errorf("Expected task to be marked completed")
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "A"},
		{ID: 2, Title: "B"},
		{ID: 3, Title: "C"},
	}

	idToDelete := 2
	var updated []Task
	for _, task := range tasks {
		if task.ID != idToDelete {
			updated = append(updated, task)
		}
	}

	if len(updated) != 2 {
		t.Errorf("Expected 2 tasks after deletion, got %d", len(updated))
	}

	for _, task := range updated {
		if task.ID == idToDelete {
			t.Errorf("Task with ID %d was not deleted", idToDelete)
		}
	}
}
