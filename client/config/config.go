package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	GameConfig  GameConfig
	LocalConfig LocalConfig
	// MultipayerClientConfig MultiplayerClientConfig
}

type LocalConfig struct {
	BallSize     int
	PlatWidth    int
	PlatHeight   int
	ScreenWidth  int
	ScreenHeight int
}

type GameConfig struct {
	MaxTPS int
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
	log.Printf("%+v\n", *conf)
}
