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
