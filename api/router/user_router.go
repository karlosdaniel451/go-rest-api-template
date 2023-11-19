package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karlosdaniel451/go-rest-api-template/api/controller"
)

func setupUserRouter(app *fiber.App, restController *controller.UserController) {
	app.Post("/users", restController.Create)
	app.Get("/users", restController.GetAll)
	app.Get("/users/:id", restController.GetById)
	app.Delete("/users/:id", restController.Delete)

	app.Post("/users/:id/tasks", restController.CreateTaskForUser)
	app.Get("/users/:id/tasks", restController.GetTasksOfUser)
}
