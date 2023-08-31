package setup

import (
	"log"

	"github.com/karlosdaniel451/go-rest-api-template/api/controller"
	"github.com/karlosdaniel451/go-rest-api-template/db"
	"github.com/karlosdaniel451/go-rest-api-template/repository"
	"github.com/karlosdaniel451/go-rest-api-template/usecase"
)

var (
	// Repositories
	TaskRepository repository.TaskRepository

	// Use cases
	TaskUseCase usecase.TaskUseCase

	// Controllers
	TaskController controller.TaskController
)

func Setup() {
	assertInterfaces()

	// Try to connect to the database server.
	err := db.Connect()
	if err != nil {
		log.Fatalf("error when connecting to database: %s", err)
	}

	// Setup for Task.
	TaskRepository = repository.NewTaskRepositoryDB(db.GetDB())
	TaskUseCase = usecase.NewTaskUseCaseImpl(TaskRepository)
	TaskController = controller.NewTaskController(TaskUseCase)
}

func assertInterfaces() {
	var _ usecase.TaskUseCase = usecase.TaskUseCaseImpl{}
	var _ repository.TaskRepository = repository.TaskRepositoryDB{}
}
