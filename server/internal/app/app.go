package app

import (
	"gopong/server/internal/game"
	"gopong/server/internal/server"
)

func Run(configPath string) {
	game := game.New()
	httpServer := server.New(game.Server)

	httpServer.Run()
}
