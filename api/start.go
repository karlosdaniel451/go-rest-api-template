package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/karlosdaniel451/go-rest-api-template/api/middleware"
	"github.com/karlosdaniel451/go-rest-api-template/api/router"
	"github.com/karlosdaniel451/go-rest-api-template/cmd/setup"
	"github.com/karlosdaniel451/go-rest-api-template/config"
)

// @title Go REST API Template
// @version 0.0.1
// @description Template for a RESTful web service in Go with Fiber.
func StartApp(config config.AppConfig) error {
	// Create a new Fiber app instance with configuration settings.
	app := fiber.New(fiber.Config{
		AppName:           "Simple Go RESTful API with Fiber and GORM",
		EnablePrintRoutes: true,
	})

	// Setup custom HTTP middlewares.
	middleware.Setup(app)

	// Setup HTTP routers with the corresponding controllers.
	router.Setup(app, &setup.TaskController, &setup.UserController)

	// Start the Fiber app and listen on the specified port.
	if err := app.Listen(fmt.Sprintf(":%d", config.ListenerPort)); err != nil {
		return err
	}

	return nil
}
