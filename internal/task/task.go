package task

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Project     string    `json:"project"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	TimeSpent   int       `json:"time_spent"`
	IsTracking  bool      `json:"is_tracking"`
	LastStarted time.Time `json:"last_started"`
	FinishedAt  time.Time `json:"finished_at"`
}

func NewTask(title string, project string, priority string) Task {
	return Task{
		Title:     title,
		Project:   project,
		Priority:  priority,
		Status:    "pending",
		CreatedAt: time.Now().UTC(),
	}
}

func (t *Task) StartTimer() {
	if !t.IsTracking {
		t.IsTracking = true
		t.LastStarted = time.Now().UTC()
	}
}

func (t *Task) StopTimer() {
	if t.IsTracking {
		t.TimeSpent += int(time.Since(t.LastStarted).Seconds())
		t.IsTracking = false
	}
}
