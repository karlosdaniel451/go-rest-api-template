package main

import (
	"fmt"
	"github.com/karlosdaniel451/go-rest-api-template/api"
	"github.com/karlosdaniel451/go-rest-api-template/cmd/setup"
	"github.com/karlosdaniel451/go-rest-api-template/config"
	_ "github.com/karlosdaniel451/go-rest-api-template/docs"
	"log/slog"
	"os"
)

func main() {
	appConfig := config.NewEmptyAppConfig()

	if err := setup.Setup(appConfig); err != nil {
		slog.Error("error when setting up application", "error", err)
		os.Exit(1)
	}

	if err := api.StartApp(*appConfig); err != nil {
		slog.Info(
			fmt.Sprintf("failed to start RESTful Web Service at %d", appConfig.ListenerPort),
			"error", err,
		)
		os.Exit(1)
	}
}
