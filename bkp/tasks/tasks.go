package tasks

import (
	"fmt"
	"time"

	"github.com/mbassini/zenwork/model"
	"github.com/mbassini/zenwork/storage"
)

// *storage.ListsData -> Access directly to the slice loaded in storage

func New(task model.Task, listID int) error {
	fmt.Printf("Adding task: %s \n", task.Name)

	found := false
	for i, l := range *storage.ListsData {
		if l.ID == listID {
			task.ID = len(l.Tasks) + 1
			task.CreatedAt = time.Now()
			(*storage.ListsData)[i].Tasks[task.ID] = task
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("could not find a list with id <%d>", listID)
	}

	return storage.SaveLists()
}

func Complete(taskID int, listID int) error {
	list, ok := (*storage.ListsData)[listID]
	if !ok {
		return fmt.Errorf("list with id <%d> not found", listID)
	}

	task, ok := list.Tasks[taskID]
	if !ok {
		return fmt.Errorf("task with id <%d> not found in list with id <%d>", taskID, listID)
	}

	task.Completed = true
	task.CompletedAt = time.Now()

	list.Tasks[taskID] = task

	(*storage.ListsData)[listID] = list

	return storage.SaveLists()
}
