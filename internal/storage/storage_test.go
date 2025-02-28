package storage

import (
	"github.com/mbassini/zenwork/internal/task"
	"github.com/stretchr/testify/assert"
	"testing"
)

// @TODO: Use table testing pattern
func TestReadWriteTasks(t *testing.T) {
	// Set a temporary dir to store test file
	t.Setenv("HOME", t.TempDir())

	tasks := []task.Task{
		task.NewTask("Task 1", "project1", "high"),
		task.NewTask("Task 2", "project2", "medium"),
	}

	err := WriteTasks(tasks)
	assert.NoError(t, err)

	loadedTasks, err := ReadTasks()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(loadedTasks))
	assert.Equal(t, 1, loadedTasks[0].ID)
	assert.Equal(t, 2, loadedTasks[1].ID)
}

func TestEmptyFile(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	tasks, err := ReadTasks()
	assert.NoError(t, err)
	assert.Empty(t, tasks)
}
