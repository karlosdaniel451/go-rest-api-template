package setup

import (
	"github.com/karlosdaniel451/go-rest-api-template/api/controller"
	"log/slog"
	"os"

	"github.com/karlosdaniel451/go-rest-api-template/db"
	"github.com/karlosdaniel451/go-rest-api-template/repository"
	"github.com/karlosdaniel451/go-rest-api-template/usecase"
)

var (
	// Logger.
	logger *slog.Logger

	// Repositories.
	TaskRepository repository.TaskRepository

	// Use cases.
	TaskUseCase usecase.TaskUseCase

	// Controllers.
	TaskController controller.TaskController
)

func Setup() {
	assertInterfaces()

	// Setup structured logger.
	setupLogger()

	// Try to connect to the database server.
	if err := db.Connect(); err != nil {
		slog.Error("error when accessing to database", "error", err)
		os.Exit(1)
	}

	slog.Info("database session created successfully")

	// Setup for Task.
	TaskRepository = repository.NewTaskRepositoryDB(db.GetDB())
	TaskUseCase = usecase.NewTaskUseCaseImpl(TaskRepository)
	TaskController = controller.NewTaskController(TaskUseCase)
}

func setupLogger() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func assertInterfaces() {
	var _ usecase.TaskUseCase = usecase.TaskUseCaseImpl{}
	var _ repository.TaskRepository = repository.TaskRepositoryDB{}
}
