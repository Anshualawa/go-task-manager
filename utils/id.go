package utils

import "go-task-manager/task"

func GetNextID(tasks []task.Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1

}
