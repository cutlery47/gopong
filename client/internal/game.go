package game

import (
	"log"

	"github.com/cutlery47/gopong/client/config"
	"github.com/cutlery47/gopong/client/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type ChannelManager struct {
	clientExitChan    chan byte
	clientFinishChan  chan byte
	clientStartChan   chan byte
	sessionExitChan   chan byte
	sessionFinishChan chan byte
	sessionStartChan  chan byte
	sessionIdleChan   chan byte
}

func (cm ChannelManager) Manage() {
	for {
		select {
		// player decides to start a new game
		case <-cm.sessionStartChan:
			// a new client is run
			cm.clientStartChan <- 1
		// player decides to quit the game
		case <-cm.sessionFinishChan:
			// both client and session are terminated
			cm.clientExitChan <- 1
			cm.sessionExitChan <- 1
		// game finishes
		case <-cm.clientFinishChan:
			// session is set back to idle
			cm.sessionIdleChan <- 1
		}
	}
}

type LocalGame struct {
	client  core.Client
	session core.Session
}

func NewLocalGame(configPath string) *LocalGame {
	config, err := config.FromFile(configPath)
	if err != nil {
		log.Fatal("rror when parsing config:", err)
	}

	state := core.StateFromConfig(config.StateConfig)

	// channel for transfering user input
	keyboardInputChan := make(chan core.CombinedKeyboardGameInputResult)

	clientExitChan := make(chan byte)
	clientFinishChan := make(chan byte)
	clientStartChan := make(chan byte)
	client := core.InitClient(keyboardInputChan, clientExitChan, clientStartChan, clientFinishChan, &state)

	sessionExitChan := make(chan byte)
	sessionFinishChan := make(chan byte)
	sessionStartChan := make(chan byte)
	sessionIdleChan := make(chan byte)
	session := core.InitSession(keyboardInputChan, sessionExitChan, sessionFinishChan, sessionStartChan, sessionIdleChan, &state)

	chanManager := ChannelManager{
		clientExitChan:    clientExitChan,
		clientFinishChan:  clientFinishChan,
		clientStartChan:   clientStartChan,
		sessionExitChan:   sessionExitChan,
		sessionFinishChan: sessionFinishChan,
		sessionStartChan:  sessionStartChan,
		sessionIdleChan:   sessionIdleChan,
	}

	game := LocalGame{
		client:  client,
		session: session,
	}

	go chanManager.Manage()

	return &game
}

func (g LocalGame) Run() {
	go g.client.Run()
	go g.session.ListenForIdle()
	ebiten.RunGame(g.session)
}

// type MultiplayerGame struct {
// 	client      core.Client
// 	session     core.Session
// 	multiplayer multiplayer.Multiplayer
// }

// func NewMultiplayerGame(configPath string) *MultiplayerGame {
// 	config, err := config.FromFile(configPath)
// 	if err != nil {
// 		log.Fatal("rror when parsing config:", err)
// 	}

// 	inputChan := make(chan core.CombinedKeyboardInputResult)
// 	outputChan := make(chan core.CombinedKeyboardInputResult)

// 	multiplayerExitChan := make(chan byte)
// 	multiplayerFinishChan := make(chan byte)
// 	multiplayer := multiplayer.Init(config, inputChan, outputChan, multiplayerExitChan, multiplayerFinishChan)

// 	state := multiplayer.FindGame()

// 	clientExitChan := make(chan byte)
// 	clientFinishChan := make(chan byte)
// 	client := core.InitClient(inputChan, clientExitChan, clientFinishChan, state)

// 	sessionExitChan := make(chan byte)
// 	session := core.InitSession(inputChan, sessionExitChan, state)

// 	game := MultiplayerGame{
// 		client:      client,
// 		session:     session,
// 		multiplayer: multiplayer,
// 	}

// 	go game.manageExitChannels(clientExitChan, clientFinishChan, sessionExitChan, multiplayerExitChan, multiplayerFinishChan)

// 	return &game
// }

// func (g MultiplayerGame) manageExitChannels(clientExit, sessionExit, multiplayerExit chan<- byte, clientFinish, multiplayerFinish <-chan byte) {
// 	for {
// 		select {
// 		case <-clientFinish:
// 			sessionExit <- 1
// 			clientExit <- 1
// 			multiplayerExit <- 1
// 		case <-multiplayerFinish:
// 			sessionExit <- 1
// 			clientExit <- 1
// 			multiplayerExit <- 1
// 		}
// 	}

// }
