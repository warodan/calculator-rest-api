package config

import (
	"github.com/go-playground/validator/v10"
	"os"
)

type Config struct {
	Port        string `validate:"required,numeric"`
	LoggerLevel string `validate:"required,oneof=DEBUG INFO WARN ERROR"`
}

func (cfg *Config) Validate() error {
	return validator.New().Struct(cfg)
}

func Load() *Config {
	config := &Config{}

	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
	} else {
		config.Port = "8080"
	}

	if level := os.Getenv("LOGGER_LEVEL"); level != "" {
		config.LoggerLevel = level
	} else {
		config.LoggerLevel = "INFO"
	}

	return config
}
