package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/chapsuk/grace"
	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/app"
	"github.com/oxyd-io/faker/internal/app/config"
	"github.com/oxyd-io/faker/internal/app/env"
	"github.com/oxyd-io/faker/internal/app/generator"
	zapLogger "github.com/oxyd-io/faker/internal/logger"
)

const (
	serviceName = "faker"
)

var (
	version = "unknown"
)

func main() {
	var environment, logLevel string

	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&environment, "e", "", "environment")
	flag.StringVar(&logLevel, "ll", "info", "logging level")
	flag.Parse()

	ctx := grace.ShutdownContext(context.Background())
	ctx = context.WithValue(ctx, env.Name, serviceName)
	ctx = context.WithValue(ctx, env.Version, version)
	ctx = context.WithValue(ctx, env.Environment, environment)

	logger, err := zapLogger.New(
		serviceName,
		version,
		environment,
		logLevel,
	)
	if err != nil {
		log.Fatal("error while init logger", zap.Error(err))
	}

	logger.Info(
		"flags",
		zap.String("version", version),
		zap.String("environment", environment),
		zap.String("log_level", logLevel),
	)

	appPath := "."
	if len(os.Getenv("APP_PATH")) > 0 {
		appPath = os.Getenv("APP_PATH")
	}

	appConfig, err := config.NewAppConfig(serviceName, appPath+"/configs/"+serviceName+"/"+environment+".yml")
	if err != nil {
		logger.Fatal("error while init config", zap.Error(err))
	}

	gnrtr := generator.New()

	application := app.New(serviceName, version, environment, appConfig, logger, gnrtr)

	application.Run(ctx)
	application.Shutdown()
}
