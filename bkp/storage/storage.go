package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mbassini/zenwork/config"
	"github.com/mbassini/zenwork/model"
)

var ListsData = &map[int]model.List{}

func LoadLists() error {
	fmt.Println("Loading Lists")

	file, err := os.Open(config.ListsFilename)
	if err != nil {
		fmt.Println("Error != nil")
		// If the file doesn't exist, initialize an empty slice.
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			emptyMap := make(map[int]model.List)
			ListsData = &emptyMap
			return nil
		}
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var data map[int]model.List
	err = json.NewDecoder(file).Decode(&data)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	ListsData = &data
	fmt.Printf("- Loaded %v Lists into memory -\n", len(data))
	return nil
}

func SaveLists() error {
	data, err := json.MarshalIndent(ListsData, "", "    ")
	if err != nil {
		return fmt.Errorf("Error serializing into JSON: %v", err)
	}

	fmt.Println(string(data))
	return os.WriteFile(config.ListsFilename, data, config.FilePermissions)
}
