package cmd

import (
	"fmt"
	"github.com/mbassini/zenwork/internal/storage"
	"github.com/spf13/cobra"
	"strconv"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Task ID is missing.")
			fmt.Println("Usage: zenwork rm [id]")
			fmt.Println("Tip: You can run ls to see current IDs")
			return
		}

		tasks, err := storage.ReadTasks()
		if err != nil {
			panic(fmt.Sprintf("Failed to read tasks: %v", err))
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Sprintf("Invalid task ID format: %v", err))
		}

		var taskIdx *int
		for i, t := range tasks {
			if t.ID == id {
				taskIdx = &i
				break
			}
		}
		if taskIdx == nil {
			fmt.Printf("Task with ID %d not found.\n", id)
			return
		}

		tasks = append(tasks[:*taskIdx], tasks[*taskIdx+1:]...)
		if err := storage.WriteTasks(tasks); err != nil {
			panic(fmt.Sprintf("Could not delete task, error trying to write file: %v", err))
		}
		fmt.Printf("Deleted task %d!\n", id)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
