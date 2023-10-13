package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Setup(app *fiber.App) {
	loggerConfig := logger.Config{
		Next: nil,
		Done: nil,
		Format: "[${time}] - ${locals:requestid} - ${ip}:${port} - ${ua} - \"${method} ${path}\" - ${status} " +
			"- ${latency}\n",

		// Format: "{\"time\": \"${time}\", \"request_id\": \"${locals:requestid}\", " +
		// 	"\"client\": \"${ip}:${port}\", \"user_agent\": \"${ua}\", " +
		// 	"\"method\": \"${method}\", \"path\": \"${path}\", \"status\": \"${status}\"}",

		// Format:        "{\"method\": \"${method}\"}",
		TimeFormat:    time.RFC3339Nano,
		TimeZone:      "UTC",
		TimeInterval:  500 * time.Millisecond,
		Output:        os.Stdout,
		DisableColors: true,
	}

	app.Use(logger.New(loggerConfig))
	app.Use(requestid.New())
}
