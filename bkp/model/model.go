package model

import "time"

type List struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	Tasks     map[int]Task `json:"tasks"`
}

type Task struct {
	Completed   bool      `json:"completed"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"due_date,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
