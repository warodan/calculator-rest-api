package config

import "os"

type Config struct {
	Port        string
	LoggerLevel string
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
