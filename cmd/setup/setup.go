package setup

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/karlosdaniel451/go-rest-api-template/api/controller"
	"github.com/karlosdaniel451/go-rest-api-template/config"
	"log/slog"
	"os"
	"strconv"

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

func Setup(appConfig *config.AppConfig) error {
	assertInterfaces()

	// Setup structured logger.
	setupLogger()

	if err := setEnvVariables(appConfig); err != nil {
		return fmt.Errorf("error when setting environment variables: %s", err)
	}

	// Try to connect to the database server.
	if err := db.Connect(*appConfig); err != nil {
		return fmt.Errorf("failed to connect to database: %s", err)
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

	return nil
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

func setEnvVariables(appConfig *config.AppConfig) error {
	// Try to load .env file for environment variables.
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error when loading .env file: %s", err)
	}

	appPortNumber, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		//slog.Error("invalid config params: invalid app port number", "error", err)
		return fmt.Errorf("invalid config params: invalid app port number: %s", err)
	}

	appEnvironmentType, err := config.ParseAppEnvironmentType(
		os.Getenv("APP_ENVIRONMENT_TYPE"),
	)
	if err != nil {
		return fmt.Errorf("invalid config params: invalid app environment")
	}

	dbPortNumber, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return fmt.Errorf("invalid config params: invalid database port number")
	}

	appConfig.ListenerPort = appPortNumber
	appConfig.AppEnvironmentType = appEnvironmentType
	appConfig.DatabaseHost = os.Getenv("DB_HOST")
	appConfig.DatabaseUser = os.Getenv("DB_USER")
	appConfig.DatabasePort = dbPortNumber
	appConfig.DatabaseName = os.Getenv("DB_NAME")
	appConfig.DatabasePassword = os.Getenv("DB_PASSWORD")

	return nil
}
