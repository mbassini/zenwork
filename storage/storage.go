package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mbassini/zenwork/config"
	"github.com/mbassini/zenwork/model"
)

var ListsData = &[]model.List{}

func LoadLists() error {
	file, err := os.Open(config.ListsFilename)
	if err != nil {
		// If the file doesn't exist, initialize an empty slice.
		if os.IsNotExist(err) {
			ListsData = &[]model.List{}
			return nil
		}
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(ListsData)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	fmt.Printf("- Loaded %v Lists into memory -\n", len(*ListsData))
	return nil
}

func WriteToFile() error {
	fmt.Printf("Writing to file: %v\n", ListsData)
	data, err := json.MarshalIndent(ListsData, "", "  ")
	if err != nil {
		return fmt.Errorf("Error serializing into JSON: %v", err)
	}

	return os.WriteFile(config.ListsFilename, data, config.FilePermissions)
}
