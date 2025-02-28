package storage

import (
	"encoding/json"
	"fmt"
	"github.com/mbassini/zenwork/internal/task"
	"os"
	"path/filepath"
)

func ReadTasks() ([]task.Task, error) {
	dataPath, err := getDataPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get data path: %w", err)
	}

	if exists := fileExists(dataPath); !exists {
		err := initFile(dataPath)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize file: %w", err)
		}
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks file: %w", err)
	}

	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks JSON: %w", err)
	}

	return tasks, nil
}

func WriteTasks(tasks []task.Task) error {
	dataPath, err := getDataPath()
	if err != nil {
		return err
	}

	if exists := fileExists(dataPath); !exists {
		err := initFile(dataPath)
		if err != nil {
			return err
		}
	}

	for i := range tasks {
		if tasks[i].ID == 0 {
			tasks[i].ID = getNextID(tasks[:i])
		}
	}

	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataPath, data, 0644)
}

func getDataPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".zenwork", "tasks.json"), nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func initFile(filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(filePath, []byte("[]"), 0644)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func getNextID(tasks []task.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}
