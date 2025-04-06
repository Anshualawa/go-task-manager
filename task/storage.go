package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const fileName = "tasks.json"

func SaveTasksToFile(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("⚠️ Failed to encode tasks :", err)
		return
	}

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("⚠️ Failed to write file :", err)
	}
}

func LoadTasksFromFile() []Task {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		// File might not exist yet - that's okay
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("⚠️ Failed to parse tasks :", err)
		return []Task{}
	}

	return tasks
}
