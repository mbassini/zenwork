package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	"github.com/google/uuid"
)

const ListsFilename = "todo_lists.zenwork.json"
const JsonFilePermissions = 0644 // RW for Owner, R for group and others

type Task struct {
	Completed   bool       `json:"completed"`
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

type List struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Tasks     []Task    `json:"tasks"`
}

type FileData []List

func listsFileExists() bool {
	_, err := os.Stat(ListsFilename)
	return err == nil
}

func addList(listName string) error {
	fmt.Printf("Creating the todo list: %s \n", ListsFilename)

	list := List{
		ID:        uuid.New().String(),
		Name:      listName,
		CreatedAt: time.Now(),
		Tasks:     []Task{},
	}

	fmt.Printf("List Type: %v \n", list)

	currentLists, err := readLists()
	if err != nil {
		return err
	}

	currentLists = append(currentLists, list)

	err = writeToFile(currentLists)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	return nil
}

func readLists() ([]List, error) {
	fmt.Println("Reading todo lists")

	// Open and read the file
	file, err := os.Open(ListsFilename)
	if err != nil {
		return []List{}, nil
		//return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	var lists []List
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&lists)
	if err != nil {
		// Return an empty slice if the file is empty or has an error
		if err.Error() == "EOF" {
			return []List{}, nil
		}
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}

	return lists, nil
}

func addTask(task Task, listID string) error {
	fmt.Printf("Adding task: %s \n", task.Name)
	currentLists, err := readLists()
	if err != nil {
		return err
	}

	listIdx := findList(listID, currentLists)
	if listIdx == -1 {
		return fmt.Errorf("Could not find a list with id <%s>", listID)
	}
	fmt.Printf("LIST: %v \n", listIdx)

	currentLists[listIdx].Tasks = append(currentLists[listIdx].Tasks, task)

	err = writeToFile(currentLists)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}

func findList(id string, lists []List) int {
	return slices.IndexFunc(lists, func(l List) bool { return l.ID == id })
}

func writeToFile(fileData FileData) error {
	data, err := json.MarshalIndent(fileData, "", "  ")
	if err != nil {
		return fmt.Errorf("Error serializing into JSON: %v", err)
	}

	fmt.Printf("JSON Data: %v \n", data)
	err = os.WriteFile(ListsFilename, data, JsonFilePermissions)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	return nil
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Println(arg)
	}

	desc := "This is the second description"
	newTask := Task{
		Completed:   false,
		ID:          uuid.New().String(),
		Name:        "Second Task",
		Description: &desc,
		DueDate:     nil,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	err := addTask(newTask, "no exist")
	if err != nil {
		log.Fatal(err)
	}

	if !listsFileExists() {
		fmt.Println("File does not exist, initializing...")
		err := addList("First List")
		if err != nil {
			log.Fatal(err)
		}

		err = addList("Second List")
		if err != nil {
			log.Fatal(err)
		}
	}
}
