package cmd

import (
	"fmt"
	"github.com/mbassini/zenwork/internal/storage"
	"github.com/mbassini/zenwork/internal/task"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var (
	filterProject string
	filterStatus  string
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.ReadTasks()
		if err != nil {
			panic(fmt.Sprintf("Failed to read tasks: %v", err))
		}

		filteredTasks := filterTasks(tasks, filterProject, filterStatus)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Project", "Priority", "Status", "Deadline"})

		for _, task := range filteredTasks {
			table.Append([]string{
				fmt.Sprintf("%d", task.ID),
				task.Title,
				task.Project,
				task.Priority,
				task.Status,
				task.Deadline.Format("2006-01-02 15:04"),
			})
		}

		table.Render()
	},
}

func filterTasks(tasks []task.Task, project, status string) []task.Task {
	var filteredTasks []task.Task
	for _, t := range tasks {
		if (project == "" || t.Project == project) && (status == "" || t.Status == status) {
			filteredTasks = append(filteredTasks, t)
		}
	}
	return filteredTasks
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&filterProject, "project", "p", "", "Filter by project")
	listCmd.Flags().StringVarP(&filterStatus, "status", "s", "", "Filter by status (pending/done)")
}
