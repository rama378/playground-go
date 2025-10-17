package config

import (
	"log"
	"os"

	"github.com/stretchr/testify/assert/yaml"
)

type Config struct {
	AppName string `yaml:"app_name"`
	Env     string `yaml:"env"`
}

func Load(path string) *Config {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read config files: %v", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return &cfg
}
