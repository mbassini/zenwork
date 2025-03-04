package cmd

import (
	"fmt"
	"strconv"

	"github.com/mbassini/zenwork/internal/storage"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [taskID]",
	Short: "Stop tracking time for a task",
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
				if !t.IsTracking {
					panic(fmt.Sprintf("Task %d is not being tracked", taskID))
				}
				tasks[i].StopTimer()
				found = true
				break
			}
		}

		if !found {
			panic(fmt.Sprintf("Task %d not found", taskID))
		}

		if err := storage.WriteTasks(tasks); err != nil {
			panic(fmt.Sprintf("Failed to save tasks: %v", err))
		}

		fmt.Printf("Stopped tracking task %d\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
