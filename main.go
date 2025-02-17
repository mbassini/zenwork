package main

import (
	"flag"
	"log"
	"time"
)

type Task struct {
	Completed   bool
	ID          int
	DueDate     time.Time
	CreatedAt   time.Time
	CompletedAt time.Time
	Name        string
	Description string
}

type List struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Tasks     []Task
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Println(arg)
	}
}
