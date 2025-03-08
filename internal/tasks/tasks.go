package tasks

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    string
	Status      string
	Deadline    time.Time
	CreatedAt   time.Time
	IsTracking  bool
	LastStarted time.Time
	TimeSpent   int
	FinishedAt  time.Time
}

type TaskOption func(*Task)

func NewTask(title string, options ...TaskOption) Task {
	t := Task{
		Title:     title,
		Priority:  "normal",
		Status:    "pending",
		CreatedAt: time.Now().UTC(),
	}

	for _, option := range options {
		option(&t)
	}

	fmt.Printf("[TASK]: %v\n", t)
	return t
}

func BuildTaskOptions(options map[string]string) []TaskOption {
	opts := []TaskOption{}

	if description, ok := options["description"]; ok {
		opts = append(opts, withDescription(description))
	}
	if status, ok := options["status"]; ok {
		opts = append(opts, withStatus(status))
	}
	if priority, ok := options["priority"]; ok {
		opts = append(opts, withPriority(priority))
	}
	return opts
}

func withDescription(description string) TaskOption {
	return func(t *Task) {
		t.Description = description
	}
}

func withStatus(status string) TaskOption {
	return func(t *Task) {
		t.Status = status
	}
}

func withPriority(priority string) TaskOption {
	return func(t *Task) {
		t.Priority = priority
	}
}
