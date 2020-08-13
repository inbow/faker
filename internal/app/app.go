package app

import (
	"context"

	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/config"
	"github.com/oxyd-io/faker/internal/server/http"
)

type (
	Application struct {
		logger *zap.Logger
		config *config.AppConfig
	}
)

func New(config *config.AppConfig, logger *zap.Logger) *Application {
	return &Application{
		config: config,
		logger: logger,
	}
}

func (app *Application) Run(ctx context.Context) {
	httpServerErrCh := http.New(ctx, app.logger, app.config)
	app.logger.Info("Server is down", zap.String("type", "http"), zap.Error(<-httpServerErrCh))
}

func (app *Application) Shutdown() {
	_ = app.logger.Sync()
}
