package game

import (
	"log"
	"strconv"

	"github.com/cutlery47/gopong/client/config"
	"github.com/cutlery47/gopong/client/internal/game/local"
	"github.com/cutlery47/gopong/client/internal/game/multiplayer"
	"github.com/cutlery47/gopong/common/conn"
	"github.com/hajimehoshi/ebiten/v2"
)

func RunLocalGame(configPath string) {
	cliConfig, err := config.FromFile(configPath)
	if err != nil {
		log.Println("Error when parsing config")
		return
	}

	client := local.NewClient(cliConfig.GameConfig)
	err = ebiten.RunGame(client)
	if err != nil {
		log.Printf("A runtime error occurred: %v", err)
	}
}

func RunMultiplayerGame(configPath string) {
	cliConfig, err := config.FromFile(configPath)
	if err != nil {
		log.Println("Error when parsing config")
		return
	}

	cliConfig.Print()

	conn, servConfig, err := conn.InitConnection(
		cliConfig.WebServerConfig.Host,
		strconv.Itoa(cliConfig.WebServerConfig.Port),
	)
	if err != nil {
		log.Println("Couldn't establish connection with the server...")
		return
	}

	client := multiplayer.NewMultiplayerClient(conn, servConfig, cliConfig.GameConfig)
	err = ebiten.RunGame(client)
	if err != nil {
		log.Printf("A runtime error occurred: %v", err)
		return
	}
}
