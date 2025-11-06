package main

import "time"

const (
	StatusTodo       = "todo"
	StatusInProgress = "inprogress"
	StatusDone       = "done"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
