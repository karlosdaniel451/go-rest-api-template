package controller

import (
	"errors"
	"github.com/karlosdaniel451/go-rest-api-template/usecase"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/errs"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
	TakUseCase  usecase.TaskUseCase
}

func NewUserController(
	userUseCase usecase.UserUseCase,
	taskUseCase usecase.TaskUseCase,
) UserController {

	return UserController{UserUseCase: userUseCase, TakUseCase: taskUseCase}
}

// Create a new User.
// @Description Create a new User.
// @Summary Create a new User.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.CreateUser true "User"
// @Success 201 {object} model.CreateUser
// @Router /users [post]
func (controller UserController) Create(c *fiber.Ctx) error {
	var newUser model.User

	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"detail": "invalid task data: " + err.Error(),
		})
	}

	newUserAllData, err := controller.UserUseCase.Create(&newUser)
	if err != nil {
		slog.Error("internal error",
			"reason", err,
			"request_id", c.Locals("requestid"),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUserAllData)
}

// Delete an User.
// @Summary Delete an User by its id.
// @Description Summary Delete an User and, in case there is no Task with the given ID,
// returns a 404 not found error.
// @Tags Users
// @Produce json
// @Param id path int true "Id of the User be deleted"
// @Success 204
// @Failure 404
// @Router /users/{id} [delete]
func (controller UserController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of user should be an integer greater than 0",
		})
	}

	err = controller.UserUseCase.DeleteById(uint(id))
	if err != nil {
		if errors.As(err, &errs.NotFoundError{}) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
		slog.Error("internal error",
			"reason", err,
			"request_id", c.Locals("requestid"),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "internal server error",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Get an User by its id.
// @Description Get an User by its id.
// @Summary Get an User by its id.
// @Tags Users
// @Produce json
// @Success 200 {object} model.User
// @Failure 404
// @Router /users/{id} [get]
func (controller UserController) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of user should be an integer greater than 0",
		})
	}

	user, err := controller.UserUseCase.GetById(uint(id))
	if err != nil {
		if errors.As(err, &errs.NotFoundError{}) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
		slog.Error("internal error",
			"reason", err,
			"request_id", c.Locals("requestid"),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// Get all Users.
// @Description Get all Users
// @Summary Get all Users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func (controller UserController) GetAll(c *fiber.Ctx) error {
	tasks, err := controller.UserUseCase.GetAll()
	if err != nil {
		slog.Error("internal error",
			"reason", err,
			"request_id", c.Locals("requestid"),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}

// Creat a new Task for an User.
// @Summary Create a new Task for an User.
// @Description Create a new Task for an User, in case there is no User with the given id,
// returns a 404 no found error.
// @Tags Users
// @Produce json
// @Param task body model.CreateTask true "Task"
// @Param user_id path int true "Id of the User for whom the Task will be created"
// @Success 201 {object} model.Task
// @Router /users/{user_id}/tasks [post]
func (controller UserController) CreateTaskForUser(c *fiber.Ctx) error {
	var newTask model.Task

	// Try to parse the user_id path param.
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of user should be an integer greater than 0",
		})
	}

	// Try to parse the Task body param.
	err = c.BodyParser(&newTask)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"detail": "invalid task data: " + err.Error(),
		})
	}

	// Check whether exist an User with the given id.
	_, err = controller.UserUseCase.GetById(uint(id))
	if err != nil {
		if errors.As(err, &errs.NotFoundError{}) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
		slog.Error("internal error",
			"reason", err,
			"request_id", c.Locals("requestid"),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "internal server error",
		})
	}

	newTask.UserId = uint(id)

	newTaskAllData, err := controller.TakUseCase.Create(&newTask)
	if err != nil {
		slog.Error("internal error",
			"reason", err,
			"request_id", c.Locals("requestid"),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newTaskAllData)
}

// Get all Tasks of an User.
// @Description Get all Tasks of an User.
// @Summary Get all Tasks of an User.
// @Tags Users
// @Produce json
// @Success 200 {array} model.Task
// @Failure 404
// @Router /users/{id}/tasks [get]
func (controller UserController) GetTasksOfUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of user should be an integer greater than 0",
		})
	}

	user, err := controller.UserUseCase.GetById(uint(id))
	if err != nil {
		if errors.As(err, &errs.NotFoundError{}) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
		slog.Error("internal error",
			"reason", err,
			"request_id", c.Locals("requestid"),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user.Tasks)
}
