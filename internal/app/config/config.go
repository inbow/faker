package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP struct {
			Host  string
			Port  int
			Check struct {
				Host string
			}
		}

		Bid struct {
			DelayMin int `mapstructure:"delay_min"`
			DelayMax int `mapstructure:"delay_max"`
		}
	}
)

func New(ct string, path string) (*Config, error) {
	viper.SetConfigType(ct)
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Config{}
	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
