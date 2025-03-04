package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mbassini/zenwork/internal/storage"
	"github.com/mbassini/zenwork/internal/task"
	"github.com/spf13/cobra"
)

var (
	exportFormat string
	exportFile   string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export tasks to CSV/JSON",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.ReadTasks()
		if err != nil {
			panic(err)
		}

		switch exportFormat {
		case "csv":
			writeCSV(tasks, exportFile)
		case "json":
			writeJSON(tasks, exportFile)
		default:
			panic("Unsupported format. Use 'csv' or 'json'")
		}
	},
}

func writeCSV(tasks []task.Task, path string) {
	file, _ := os.Create(path)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{
		"ID",
		"Title",
		"Project",
		"Priority",
		"Status",
		"Time Spent (h)",
	})

	for _, t := range tasks {
		writer.Write([]string{
			fmt.Sprintf("%d", t.ID),
			t.Title,
			t.Project,
			t.Priority,
			t.Status,
			fmt.Sprintf("%.1f", float64(t.TimeSpent/3600)),
		})
	}

	writer.Flush()
}

func writeJSON(tasks []task.Task, path string) {
	data, _ := json.MarshalIndent(tasks, "", "    ")
	os.WriteFile(path, data, 0644)
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&exportFormat, "format", "f", "csv", "Export format (csv/json)")
	exportCmd.Flags().StringVarP(&exportFile, "output", "o", "tasks.csv", "Output file name")
}
