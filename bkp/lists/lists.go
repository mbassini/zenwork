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

	newID := len(*storage.ListsData) + 1
	newList := model.List{
		ID:        newID,
		Name:      name,
		CreatedAt: time.Now(),
		Tasks:     map[int]model.Task{},
	}

	(*storage.ListsData)[newID] = newList
	return storage.SaveLists()
}

func Find(id int) (model.List, error) {
	if list, ok := (*storage.ListsData)[id]; ok {
		return list, nil
	}
	return model.List{}, fmt.Errorf("list with ID <%d> not found", id)
}

func GetAll() []model.List {
	listsSlice := make([]model.List, 0, len(*storage.ListsData))
	for _, list := range *storage.ListsData {
		listsSlice = append(listsSlice, list)
	}
	return listsSlice
}
