package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	GameConfig      GameConfig
	WebServerConfig WebServerConfig
}

type GameConfig struct {
	CanvasConfig     CanvasConfig
	PlatformConfig   PlatformConfig
	BallConfig       BallConfig
	GameServerConfig GameServerConfig
}

type CanvasConfig struct {
	Width  int
	Height int
}

type PlatformConfig struct {
	Width  int
	Height int
	Speed  int
}

type BallConfig struct {
	Size int
}

type GameServerConfig struct {
	Tickrate int
}

type WebServerConfig struct {
	Port int
	Host string
}

func FromFile(configPath string) (Config, error) {
	config := Config{}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (conf *Config) Print() {
	log.Printf("%v+\n", *conf)
}
