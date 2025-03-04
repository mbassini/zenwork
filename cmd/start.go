package cmd

import (
	"fmt"
	"strconv"

	"github.com/mbassini/zenwork/internal/storage"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [taskID]",
	Short: "Start tracking time for a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			panic("Invalid task ID. Must be a number")
		}

		tasks, err := storage.ReadTasks()
		if err != nil {
			panic(fmt.Sprintf("Failed to read tasks: %v", err))
		}

		found := false
		for i, t := range tasks {
			if t.ID == taskID {
				if t.IsTracking {
					panic(fmt.Sprintf("Task %d is already being tracked", taskID))
				}

				for _, otherTask := range tasks {
					if otherTask.IsTracking {
						panic(fmt.Sprintf("Task %d is already being tracked. Stop it first.", otherTask.ID))
					}
				}
				tasks[i].StartTimer()
				found = true
				break
			}
		}

		if !found {
			panic(fmt.Sprintf("Task %d not found", taskID))
		}

		if err := storage.WriteTasks(tasks); err != nil {
			panic(fmt.Sprintf("Failed to save task %d\n", taskID))
		}

		fmt.Printf("Started tracking task %d\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
