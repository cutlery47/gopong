package game

import (
	"log"

	"github.com/cutlery47/gopong/client/config"
	"github.com/cutlery47/gopong/client/internal/client/local"
	"github.com/cutlery47/gopong/client/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Client interface {
	Run()
}

type Game struct {
	client  Client
	session core.Session
}

func NewLocalGame(configPath string) *Game {
	config, err := config.FromFile(configPath)
	if err != nil {
		log.Fatal("rror when parsing config:", err)
	}

	inputChan := make(chan core.CombinedKeyboardInputResult)
	updateChan := make(chan core.StateUpdate)

	client := local.InitLocalClient(updateChan, inputChan, config.LocalConfig)
	session := core.InitSession(updateChan, inputChan, config.LocalConfig, config.GameConfig.MaxTPS)

	return &Game{
		client:  client,
		session: session,
	}
}

// func NewMultiplayerGame(configPath string) *Game {
// 	config, err := config.FromFile(configPath)
// 	if err != nil {
// 		log.Fatal("rror when parsing config:", err)
// 	}

// 	inputChan := make(chan core.KeyboardInputResult)
// 	updateChan := make(chan core.StateUpdate)

// 	return &Game{
// 		client:  local.InitMultiplayerClient(updateChan, inputChan, config.ClientConfig),
// 		session: core.InitSession(updateChan, inputChan, config.SessionConfig),
// 	}
// }

func (g Game) Run() {
	go g.client.Run()
	ebiten.RunGame(g.session)
}

// func RunLocalGame(configPath string) {
// 	// reading config
// 	cliConfig, err := config.FromFile(configPath)
// 	if err != nil {
// 		log.Println("Error when parsing config")
// 		return
// 	}

// 	// booting up a client
// 	client := local.NewClient(cliConfig.GameConfig)
// 	err = ebiten.RunGame(client)
// 	if err != nil {
// 		log.Printf("A runtime error occurred: %v", err)
// 	}
// }

// func RunMultiplayerGame(configPath string) {
// 	// reading config
// 	cliConfig, err := config.FromFile(configPath)
// 	if err != nil {
// 		log.Println("Error when parsing config")
// 		return
// 	}

// 	// connection to the server and waiting for a game
// 	conn, servConfig, err := conn.InitClientConnection(
// 		cliConfig.WebServerConfig.Host,
// 		strconv.Itoa(cliConfig.WebServerConfig.Port),
// 	)
// 	if err != nil {
// 		log.Println("Couldn't establish connection with the server...")
// 		return
// 	}

// 	// booting up a client
// 	client := multiplayer.NewMultiplayerClient(conn, servConfig, cliConfig.GameConfig)
// 	err = ebiten.RunGame(client)
// 	if err != nil {
// 		log.Printf("A runtime error occurred: %v", err)
// 		return
// 	}
// }
