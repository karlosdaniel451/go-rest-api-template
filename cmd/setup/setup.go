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
	UserRepository repository.UserRepository

	// Use cases.
	TaskUseCase usecase.TaskUseCase
	UserUseCase usecase.UserUseCase

	// Controllers.
	TaskController controller.TaskController
	UserController controller.UserController
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

	// Setup for User.
	UserRepository = repository.NewUserRepositoryDB(db.GetDB())
	UserUseCase = usecase.NewUserUseCaseImpl(UserRepository)
	UserController = controller.NewUserController(UserUseCase, TaskUseCase)
}

func setupLogger() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func assertInterfaces() {
	// Assertions for Task.
	var _ usecase.TaskUseCase = usecase.TaskUseCaseImpl{}
	var _ repository.TaskRepository = repository.TaskRepositoryDB{}

	// Assertions for User.
	var _ usecase.UserUseCase = usecase.UserUseCaseImpl{}
	var _ repository.UserRepository = repository.UserRepositoryDB{}
}
