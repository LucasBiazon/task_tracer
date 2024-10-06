package usecases

import (
	"errors"

	entity "github.com/lucasBiazon/task_tracker/internal/entities"
	"github.com/lucasBiazon/task_tracker/internal/repository"
)

type TaskCases struct {
	repo repository.TaskRepository
}

func NewTaskCases(repo repository.TaskRepository) *TaskCases {
	return &TaskCases{repo: repo}
}

func (tc *TaskCases) CreateTask(description string) error {
	if description == "" {
		return errors.New("description is required")
	}
	return tc.repo.Create(description)
}

func (tc *TaskCases) GetTasks() ([]*entity.Task, error) {
	return tc.repo.FindAll()
}

func (tc *TaskCases) GetTask(id string) (*entity.Task, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return tc.repo.FindByID(id)
}

func (tc *TaskCases) UpdateTask(id string, description string) error {
	if description == "" {
		return errors.New("description is required")
	}
	if id == "" {
		return errors.New("id is required")
	}
	return tc.repo.Update(id, description)
}

func (tc *TaskCases) DeleteTask(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return tc.repo.Delete(id)
}

func (tc *TaskCases) CompleteTask(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return tc.repo.MarkAsDone(id)
}

func (tc *TaskCases) InProgressTask(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return tc.repo.MarkAsInProgress(id)
}

func (tc *TaskCases) GetDoneTasks() ([]*entity.Task, error) {
	return tc.repo.FindAllDone()
}

func (tc *TaskCases) GetInProgressTasks() ([]*entity.Task, error) {
	return tc.repo.FindAllInProgress()
}

func (tc *TaskCases) GetTodoTasks() ([]*entity.Task, error) {
	return tc.repo.FindAllTodo()
}
