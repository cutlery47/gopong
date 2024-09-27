package app

import (
	"log"

	"github.com/cutlery47/gopong/server/config"
	"github.com/cutlery47/gopong/server/internal/game"
	"github.com/cutlery47/gopong/server/internal/server"
)

func Run(configPath string) {
	config, err := config.FromFile(configPath)
	if err != nil {
		log.Println("Error when parsing config:", err)
		return
	}

	config.Print()

	host := config.WebServerConfig.Host
	port := config.WebServerConfig.Port

	game := game.New(config.GameConfig)

	httpServer := server.New(game.Server, server.HostPortAddr(host, port))
	httpServer.Run()
}
