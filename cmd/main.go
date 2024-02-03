// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"todo_clean_architecture/entities"
	ports "todo_clean_architecture/ports/http"
	"todo_clean_architecture/repositories"
	"todo_clean_architecture/usecases"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	db, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Auto Migrate the schema
	dbErr := db.AutoMigrate(&entities.Task{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	taskRepository := repositories.NewTaskRepository(db)
	taskService := usecases.NewTaskService(taskRepository)
	taskHandler := ports.NewTaskHandler(taskService)

	router := mux.NewRouter()

	taskHandler.RegisterHandlers(router)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}
