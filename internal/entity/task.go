package entity

type Task struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status string `json:"status" db:"status"` // "todo", "in-progress", "done"
}
