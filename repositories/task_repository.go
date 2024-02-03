package repositories

import (
	"errors"
	"todo_clean_architecture/entities"

	"github.com/jinzhu/gorm"
)

// DB defines the database methods used by the application
type ITaskRepository interface {
	GetAll() ([]entities.Task, error)
	GetByID(id uint) (*entities.Task, error)
	Create(task *entities.Task) error
	Update(task *entities.Task) error
	Delete(id uint) error
}

// db implements DB
type TaskRepository struct {
	TaskDB *gorm.DB
}

// NewDB creates a new DB
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{TaskDB: db}
}

// GetAll returns all tasks
func (d *TaskRepository) GetAll() ([]entities.Task, error) {
	var tasks []entities.Task
	result := d.TaskDB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// GetByID returns a task by ID
func (d *TaskRepository) GetByID(id uint) (*entities.Task, error) {
	var task entities.Task
	result := d.TaskDB.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &task, nil
}

// Create creates a new task
func (d *TaskRepository) Create(task *entities.Task) error {
	result := d.TaskDB.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates a task
func (d *TaskRepository) Update(task *entities.Task) error {
	result := d.TaskDB.Save(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete deletes a task by ID
func (d *TaskRepository) Delete(id uint) error {
	result := d.TaskDB.Delete(&entities.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
