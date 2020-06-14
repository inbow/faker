package logger

import (
	"go.uber.org/zap"
)

func New(service, version, environment, level string) (*zap.Logger, error) {
	var config zap.Config
	if environment == "local" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	err := config.Level.UnmarshalText([]byte(level))
	if err != nil || len(level) == 0 {
		config.Level.SetLevel(zap.DebugLevel)
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	logger = logger.With(
		zap.String("service", service),
		zap.String("version", version),
		zap.String("environment", environment),
	)

	return logger, nil
}
