package repository

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	entity "github.com/lucasBiazon/task_tracker/internal/entities"
)

type TaskRepository struct {
	filePath string
}

func NewTaskRepository(filePath string) *TaskRepository {
	return &TaskRepository{filePath: filePath}
}

func (r *TaskRepository) loadTasks() ([]*entity.Task, error) {
	file, err := os.OpenFile(r.filePath, os.O_RDONLY|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tasks := []*entity.Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) saveTasks(tasks []*entity.Task) error {
	file, err := os.OpenFile(r.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

func (r *TaskRepository) Create(description string) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}
	task := entity.NewTask(description)
	tasks = append(tasks, task)
	return r.saveTasks(tasks)
}

func (r *TaskRepository) FindAll() ([]*entity.Task, error) {
	return r.loadTasks()
}

func (r *TaskRepository) FindAllDone() ([]*entity.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}
	doneTasks := []*entity.Task{}
	for _, task := range tasks {
		if task.Status == entity.StatusDone {
			doneTasks = append(doneTasks, task)
		}
	}
	return doneTasks, nil
}

func (r *TaskRepository) FindAllInProgress() ([]*entity.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}
	inProgressTasks := []*entity.Task{}
	for _, task := range tasks {
		if task.Status == entity.StatusInProgress {
			inProgressTasks = append(inProgressTasks, task)
		}
	}
	return inProgressTasks, nil
}

func (r *TaskRepository) FindAllTodo() ([]*entity.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}
	todoTasks := []*entity.Task{}
	for _, task := range tasks {
		if task.Status == entity.StatusTodo {
			todoTasks = append(todoTasks, task)
		}
	}
	return todoTasks, nil
}

func (r *TaskRepository) FindByID(id string) (*entity.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, errors.New("tarefa não encontrada")
}

func (r *TaskRepository) Update(id, description string) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task.ID == id {
			task.Description = description
			task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return r.saveTasks(tasks)
		}
	}
	return errors.New("tarefa não encontrada")
}

func (r *TaskRepository) Delete(id string) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return r.saveTasks(tasks)
		}
	}
	return errors.New("tarefa não encontrada")
}

func (r *TaskRepository) markAsStatus(id, status string) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task.ID == id {
			task.Status = status
			task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return r.saveTasks(tasks)
		}
	}
	return errors.New("tarefa não encontrada")
}

func (r *TaskRepository) MarkAsDone(id string) error {
	err := r.markAsStatus(id, entity.StatusDone)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) MarkAsInProgress(id string) error {
	err := r.markAsStatus(id, entity.StatusInProgress)
	if err != nil {
		return err
	}
	return nil
}
