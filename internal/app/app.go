package app

import (
	"context"

	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/app/config"
	"github.com/oxyd-io/faker/internal/app/generator"
	"github.com/oxyd-io/faker/internal/app/http/server"
)

type (
	Application struct {
		Name        string
		Version     string
		Environment string

		logger *zap.Logger
		config *config.AppConfig

		generator generator.IGenerator
	}
)

func New(config *config.AppConfig, logger *zap.Logger, generator generator.IGenerator) *Application {
	return &Application{
		config: config,
		logger: logger,

		generator: generator,
	}
}

func (app *Application) Run(ctx context.Context) {
	httpServerErrCh := server.New(ctx, app.logger, app.config, app.generator)

	app.logger.Info("Server is down", zap.String("type", "http"), zap.Error(<-httpServerErrCh))
}

func (app *Application) Shutdown() {
	_ = app.logger.Sync()
}
