package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Window `yaml:"window"`
	}

	Window struct {
		Height int `yaml:"height"`
		Width  int `yaml:"width"`
	}
)

func New(filePath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(filePath, cfg)
	if err != nil {
		return cfg, fmt.Errorf("ReadConfig: %w", err)
	}
	return cfg, nil
}
