package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	GameConfig   GameConfig
	StateConfig  StateConfig
	CanvasConfig CanvasConfig
	// MultipayerClientConfig MultiplayerClientConfig
}

type StateConfig struct {
	BallSize      float64
	BallInitVelX  float64
	BallInitVelY  float64
	BallAccelMult float64
	PlatWidth     float64
	PlatHeight    float64
	PlatVelocity  float64
	ScreenWidth   float64
	ScreenHeight  float64
	PointsToWin   int
}

type CanvasConfig struct {
	ScoreX float64
	ScoreY float64
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
