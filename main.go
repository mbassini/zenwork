package main

import (
	"flag"
	"log"
	"time"
)

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
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Tasks     []Task    `json:"tasks"`
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Println(arg)
	}
}
