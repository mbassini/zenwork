package cmd

import (
	"fmt"
	"github.com/mbassini/zenwork/internal/storage"
	"github.com/mbassini/zenwork/internal/task"
	"github.com/spf13/cobra"
	"time"
)

var (
	project  string
	priority string
	deadline string
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		existingTasks, err := storage.ReadTasks()
		if err != nil {
			panic(fmt.Sprintf("Failed to read tasks: %v", err))
		}

		newTask := task.NewTask(args[0], project, priority)

		// @TODO: Do not want to store a time.Time zero value in JSON file
		if deadline != "" {
			parsedDeadline := parseDeadline(deadline)
			newTask.Deadline = parsedDeadline
		}

		existingTasks = append(existingTasks, newTask)

		if err := storage.WriteTasks(existingTasks); err != nil {
			panic(fmt.Sprintf("Failed to save task: %v", err))
		}

		fmt.Println("Task added!")
	},
}

func parseDeadline(deadline string) time.Time {
	switch deadline {
	case "today":
		// @TODO: Calc remaining time until the end of day
		return time.Now().UTC().Add(8 * time.Hour)
	case "tomorrow":
		return time.Now().UTC().AddDate(0, 0, 1)
	default:
		parsedTime, err := time.Parse(time.RFC3339, deadline)
		if err != nil {
			panic(fmt.Sprintf("Invalid deadline format. Use RFC3339 (e.g, 2024-05-16T15:00:00Z: %v", err))
		}
		return parsedTime
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&project, "project", "p", "", "Project name")
	addCmd.Flags().StringVarP(&priority, "priority", "P", "", "Priority (low/medium/high)")
	addCmd.Flags().StringVarP(&deadline, "deadline", "d", "", "Deadline (today/tomorrow/RFC3339)")
}
