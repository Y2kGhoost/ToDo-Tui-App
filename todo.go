package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const (
	dataDir  = "data"
	fileName = "tasks.json"
)

func SaveTasks(tasks []Task) error {
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(dataDir, fileName)

	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	// 4. Write the file
	return os.WriteFile(fullPath, data, 0644)
}

func LoadTasks() ([]Task, error) {
	fullPath := filepath.Join(dataDir, fileName)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
