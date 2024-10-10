package game

import (
	"log"

	"github.com/cutlery47/gopong/client/config"
	"github.com/cutlery47/gopong/client/internal/core"
	"github.com/cutlery47/gopong/client/internal/multiplayer"
	"github.com/hajimehoshi/ebiten/v2"
)

type LocalGame struct {
	client  core.Client
	session core.Session
}

func NewLocalGame(configPath string) *LocalGame {
	config, err := config.FromFile(configPath)
	if err != nil {
		log.Fatal("rror when parsing config:", err)
	}

	keyboardInputChan := make(chan core.CombinedKeyboardInputResult)
	clientExitChan := make(chan byte)
	clientFinishChan := make(chan byte)
	sessionExitChan := make(chan byte)

	state := core.StateFromConfig(config.StateConfig)
	client := core.InitClient(keyboardInputChan, clientExitChan, clientFinishChan, &state)
	session := core.InitSession(keyboardInputChan, sessionExitChan, &state)

	game := LocalGame{
		client:  client,
		session: session,
	}

	go game.manageExitChannels(clientExitChan, clientFinishChan, sessionExitChan)

	return &game
}

func (g LocalGame) Run() {
	go g.client.Run()
	ebiten.RunGame(g.session)
}

func (g LocalGame) manageExitChannels(clientExit, clientFinish, sessionExit chan byte) {
	<-clientFinish
	sessionExit <- 1
	clientExit <- 1
}

type MultiplayerGame struct {
	client      core.Client
	session     core.Session
	multiplayer multiplayer.Multiplayer
}

func NewMultiplayerGame(configPath string) *MultiplayerGame {
	config, err := config.FromFile(configPath)
	if err != nil {
		log.Fatal("rror when parsing config:", err)
	}

	inputChan := make(chan core.CombinedKeyboardInputResult)
	outputChan := make(chan core.CombinedKeyboardInputResult)

	multiplayerExitChan := make(chan byte)
	multiplayerFinishChan := make(chan byte)
	multiplayer := multiplayer.Init(config, inputChan, outputChan, multiplayerExitChan, multiplayerFinishChan)

	state := multiplayer.FindGame()

	clientExitChan := make(chan byte)
	clientFinishChan := make(chan byte)
	client := core.InitClient(inputChan, clientExitChan, clientFinishChan, state)

	sessionExitChan := make(chan byte)
	session := core.InitSession(inputChan, sessionExitChan, state)

	game := MultiplayerGame{
		client:      client,
		session:     session,
		multiplayer: multiplayer,
	}

	go game.manageExitChannels(clientExitChan, clientFinishChan, sessionExitChan, multiplayerExitChan, multiplayerFinishChan)

	return &game
}

func (g MultiplayerGame) manageExitChannels(clientExit, sessionExit, multiplayerExit chan<- byte, clientFinish, multiplayerFinish <-chan byte) {
	for {
		select {
		case <-clientFinish:
			sessionExit <- 1
			clientExit <- 1
			multiplayerExit <- 1
		case <-multiplayerFinish:
			sessionExit <- 1
			clientExit <- 1
			multiplayerExit <- 1
		}
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
