package app

import (
	"context"

	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/config"
	"github.com/oxyd-io/faker/internal/http"
)

type (
	Application struct {
		logger *zap.Logger
		config *config.AppConfig

		httpServer *http.Server
	}
)

func New(config *config.AppConfig, logger *zap.Logger, httpServer *http.Server) *Application {
	return &Application{
		config: config,
		logger: logger,

		httpServer: httpServer,
	}
}

func (app *Application) Run(ctx context.Context) {
	go func() {
		if err := app.httpServer.Start(); err != nil {
			app.logger.Error("error while running http server", zap.Error(err))
		}
	}()

	<-ctx.Done()
}

func (app *Application) Shutdown() {
	app.logger.Info("Stop http server")
	app.httpServer.Stop() // nolint:errcheck

	app.logger.Info("Shutdown logger")
	app.logger.Sync() // nolint:errcheck
}
