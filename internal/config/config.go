package config

import (
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		HTTP struct {
			Host string `mapstructure:"host"`
			Port int    `mapstructure:"port"`

			HealthCheck struct {
				Host string `mapstructure:"host"`
			} `mapstructure:"check"`
		}

		Bid struct {
			DelayMin int `mapstructure:"delay_min"`
			DelayMax int `mapstructure:"delay_max"`
		}
	}
)

func New(serviceName, configFile string) (*AppConfig, error) {
	config, err := loadConfig(serviceName, configFile)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfig(serviceName, configFile string) (*AppConfig, error) {
	var appConfig AppConfig

	viper.AutomaticEnv()
	viper.SetEnvPrefix(serviceName)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, err
}
