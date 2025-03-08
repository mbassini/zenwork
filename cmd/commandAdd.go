package cmd

import (
	"fmt"

	"github.com/mbassini/zenwork/internal/tasks"
)

func commandAdd(positionArg string, options map[string]string) error {
	title := positionArg

	taskOptions := tasks.BuildTaskOptions(options)
	task := tasks.NewTask(title, taskOptions...)
	fmt.Printf("Task %v created!\n", task.ID)
	return nil
}
