package service

import (
	"github.com/KONICCO/go-industrious-boilerplate/internal/model"
	"github.com/KONICCO/go-industrious-boilerplate/internal/repository"
)

type TaskService interface {
	CreateTask(task *model.Task) error
	GetTask(id int) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
	ListTasks() ([]*model.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func New(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *model.Task) error {
	return s.repo.Create(task)
}

func (s *taskService) GetTask(id int) (*model.Task, error) {
	return s.repo.Get(id)
}

func (s *taskService) UpdateTask(task *model.Task) error {
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}

func (s *taskService) ListTasks() ([]*model.Task, error) {
	return s.repo.List()
}
