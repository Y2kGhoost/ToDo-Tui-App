package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const TodoFile = "./data/tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(TodoFile); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(TodoFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(TodoFile, data, 0644)
}
