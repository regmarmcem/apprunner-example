package models

import "time"

type Task struct {
	ID        int       `json:"task_id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
}
