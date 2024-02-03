// ports/http/task_handler.go
package ports

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo_clean_architecture/entities"
	"todo_clean_architecture/usecases"

	"github.com/gorilla/mux"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	taskService usecases.TaskService
}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler(taskService usecases.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// RegisterHandlers registers HTTP handlers for tasks
func (h *TaskHandler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/tasks", h.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", h.GetTask).Methods("GET")
	router.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")
}

// GetTasks handles GET request for all tasks
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskService.GetTasks()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	respondJSON(w, tasks)
}

// GetTask handles GET request for a single task by ID
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTaskByID(uint(id))
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	respondJSON(w, task)
}

// CreateTask handles POST request to create a new task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask entities.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.taskService.CreateTask(&newTask)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	respondJSON(w, newTask)
}

// UpdateTask handles PUT request to update a task by ID
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask entities.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedTask.ID = uint(id)
	err = h.taskService.UpdateTask(&updatedTask)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	respondJSON(w, updatedTask)
}

// DeleteTask handles DELETE request to delete a task by ID
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.taskService.DeleteTask(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
