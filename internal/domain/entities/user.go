package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatAt     string `json:"creat_at"`
	UpdateAt    string `json:"update_at"`
}

func NewUser(description string) *User {
	return &User{
		ID:          uuid.New().String(),
		Description: description,
		Status:      "todo",
		CreatAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
}

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]*User, error)
	FindByID(id string) (*User, error)
	Update(id, description string) error
	Delete(id string) error
	MarkAsDone(id string) error
	MarkAsInProgress(id string) error
}
