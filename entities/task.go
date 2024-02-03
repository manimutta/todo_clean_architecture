// entities/task.go
package entities

import "gorm.io/gorm"

// Task represents a task in the TODO app
type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Notes     string `json:"notes"`
}
