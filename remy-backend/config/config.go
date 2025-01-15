package config

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Server      struct {
		Host string
		Port int
	}
}

func LoadConfiguration(configPath, configName, configType string) (*Configuration, error) {
	var config *Configuration

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		slog.Info("could not read the config file: %v", err)
		return nil, fmt.Errorf("could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		slog.Info("could not unmarshal: %v", err)
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	return config, nil
}
