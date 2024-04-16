package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	API API
}

type API struct {
	Host         string
	Port         int
	WriteTimeout int
	ReadTimeout  int
	IdleTimeout  int
}

func NewConfig(filePath string) (*Config, error) {
	config, err := readConfig(filePath)
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		API: API{
			Host:         config.GetString("api.http.host"),
			Port:         config.GetInt("api.http.port"),
			WriteTimeout: config.GetInt("api.http.writeTimeout"),
			ReadTimeout:  config.GetInt("api.http.readTimeout"),
			IdleTimeout:  config.GetInt("api.http.idleTimeout"),
		},
	}, nil
}

func readConfig(filePath string) (*viper.Viper, error) {
	config := viper.New()
	config.SetEnvPrefix("strava_analytics_api")
	config.SetConfigType("yaml")
	config.SetConfigFile(filePath)
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("readConfig - failed to read config file using viper: %w", err)
	}

	return config, nil
}
