package config

import (
	"os"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Port() string {
	return c.getValue("PORT", "4500")
}

func (c *Config) getValue(envName string, defaultValue string) string {
	if os.Getenv(envName) != "" {
		return os.Getenv(envName)
	}
	return defaultValue
}
