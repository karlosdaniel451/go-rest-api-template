package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/karlosdaniel451/go-rest-api-template/api/middleware"
	"github.com/karlosdaniel451/go-rest-api-template/api/router"
	"github.com/karlosdaniel451/go-rest-api-template/cmd/setup"
	_ "github.com/karlosdaniel451/go-rest-api-template/docs"
)

var port = os.Getenv("API_PORT")

// @title Go REST API Template
// @version 0.0.1
// @description Template for a RESTful web service in Go with Fiber.
func main() {
	setup.Setup()

	app := fiber.New(fiber.Config{
		AppName:           "Simple Go RESTful API with Fiber and GORM",
		EnablePrintRoutes: true,
	})

	middleware.Setup(app)

	router.Setup(app, &setup.TaskController)

	log.Fatal(app.Listen(":" + port))
}
