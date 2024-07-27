package model

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` // "todo", "in-progress", "done"
}
