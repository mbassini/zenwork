package lists

import (
	"fmt"
	"time"

	"github.com/mbassini/zenwork/model"
	"github.com/mbassini/zenwork/storage"
)

// *storage.ListsData -> Access directly to the slice loaded in storage

func New(name string) error {
	fmt.Printf("Creating list: %s \n", name)

	for _, l := range *storage.ListsData {
		if l.Name == name {
			return fmt.Errorf("A list with name <%s> already exists.", name)
		}
	}

	newList := model.List{
		ID:        len(*storage.ListsData) + 1,
		Name:      name,
		CreatedAt: time.Now(),
		Tasks:     []model.Task{},
	}

	*storage.ListsData = append(*storage.ListsData, newList)
	return storage.SaveLists()
}

func Find(id int) (model.List, error) {
	for _, l := range *storage.ListsData {
		if l.ID == id {
			return l, nil
		}
	}
	return model.List{}, fmt.Errorf("List with ID <%d> not found.", id)
}
