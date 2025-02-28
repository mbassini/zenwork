package main

import (
	"fmt"
	"github.com/mbassini/zenwork/internal/storage"
	"github.com/mbassini/zenwork/internal/task"
	"time"
)

func main() {
	// Read current tasks
	existingTasks, err := storage.ReadTasks()
	if err != nil {
		panic(fmt.Sprintf("Failed to read tasks: %v", err))
	}

	// Create and append new task
	newTask := task.NewTask("Other test", "Testing", "low")
	newTask.Deadline = time.Now().UTC().Add(12 * time.Hour)

	existingTasks = append(existingTasks, newTask)

	// Save all tasks
	if err := storage.WriteTasks(existingTasks); err != nil {
		panic(fmt.Sprintf("Failed to write tasks: %v", err))
	}

	// Read back
	loadedTasks, err := storage.ReadTasks()
	if err != nil {
		panic(fmt.Sprintf("Failed to reload tasks: %v", err))
	}

	fmt.Printf("Loaded tasks: %+v\n", loadedTasks)
}
