package repository

import (
	"database/sql"

	"github.com/KONICCO/go-industrious-boilerplate/internal/model"
)

type TaskRepository interface {
	Create(task *model.Task) error
	Get(id int) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
	List() ([]*model.Task, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) TaskRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(task *model.Task) error {
	query := `INSERT INTO tasks (title, status) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, task.Title, task.Status).Scan(&task.ID)
}

func (r *postgresRepository) Get(id int) (*model.Task, error) {
	task := &model.Task{}
	query := `SELECT id, title, status FROM tasks WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Status)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *postgresRepository) Update(task *model.Task) error {
	query := `UPDATE tasks SET title = $1, status = $2 WHERE id = $3`
	_, err := r.db.Exec(query, task.Title, task.Status, task.ID)
	return err
}

func (r *postgresRepository) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *postgresRepository) List() ([]*model.Task, error) {
	query := `SELECT id, title, status FROM tasks`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		task := &model.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
