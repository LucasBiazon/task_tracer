package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "inprogress"
	StatusDone       = "done"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewTask(description string) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}

type TaskRepository interface {
	Create(Task *Task) error
	FindAll() ([]*Task, error)
	FindAllDone() ([]*Task, error)
	FindAllInProgress() ([]*Task, error)
	FindAllTodo() ([]*Task, error)
	FindByID(id string) (*Task, error)
	Update(id, description string) error
	Delete(id string) error
	MarkAsDone(id string) error
	MarkAsInProgress(id string) error
}
