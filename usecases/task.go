// usecases/task.go
package usecases

import (
	"todo_clean_architecture/entities"
	"todo_clean_architecture/repositories"
)

// TaskService provides use cases for tasks
type TaskService interface {
	GetTasks() ([]entities.Task, error)
	GetTaskByID(id uint) (*entities.Task, error)
	CreateTask(task *entities.Task) error
	UpdateTask(task *entities.Task) error
	DeleteTask(id uint) error
}

// taskService implements TaskService
type taskService struct {
	taskRepository repositories.ITaskRepository
}

// NewTaskService creates a new TaskService
func NewTaskService(taskRepository repositories.ITaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

// GetTasks returns all tasks
func (s *taskService) GetTasks() ([]entities.Task, error) {
	return s.taskRepository.GetAll()
}

// GetTaskByID returns a task by ID
func (s *taskService) GetTaskByID(id uint) (*entities.Task, error) {
	return s.taskRepository.GetByID(id)
}

// CreateTask creates a new task
func (s *taskService) CreateTask(task *entities.Task) error {
	return s.taskRepository.Create(task)
}

// UpdateTask updates a task
func (s *taskService) UpdateTask(task *entities.Task) error {
	return s.taskRepository.Update(task)
}

// DeleteTask deletes a task by ID
func (s *taskService) DeleteTask(id uint) error {
	return s.taskRepository.Delete(id)
}
