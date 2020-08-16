package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/chapsuk/grace"
	ll "github.com/oxyd-io/go-logger"
	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/app"
	"github.com/oxyd-io/faker/internal/config"
	"github.com/oxyd-io/faker/pkg/env"
)

const (
	service = "faker"
)

var (
	version = "unknown"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := grace.ShutdownContext(context.Background())

	var environment, logLevel string
	logLevel = os.Getenv("LOG_LEVEL")
	environment = os.Getenv("ENVIRONMENT")

	ctx = context.WithValue(ctx, env.Name, service)
	ctx = context.WithValue(ctx, env.Version, version)
	ctx = context.WithValue(ctx, env.LogLevel, logLevel)
	ctx = context.WithValue(ctx, env.Environment, environment)

	logger, err := ll.New(service, version, environment, logLevel)
	if err != nil {
		log.Fatal("error while init logger", zap.Error(err))
	}

	appPath := "."
	if len(os.Getenv("APP_PATH")) > 0 {
		appPath = os.Getenv("APP_PATH")
	}

	appConfig, err := config.New(service, appPath+"/configs/"+service+"/"+environment+".yml")
	if err != nil {
		logger.Fatal("error while init config", zap.Error(err))
	}

	hostname, _ := os.Hostname()
	appConfig.HTTP.Host = hostname

	application := app.New(appConfig, logger)

	application.Run(ctx)
	application.Shutdown()
}
