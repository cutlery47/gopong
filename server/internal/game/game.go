package game

import (
	"github.com/cutlery47/gopong/server/config"

	"github.com/cutlery47/gopong/server/internal/game/queue"
	"github.com/cutlery47/gopong/server/internal/game/server"
	"github.com/cutlery47/gopong/server/internal/game/session"

	"github.com/cutlery47/gopong/common/conn"
)

// the game itself
type Game struct {
	Server         *server.Server
	Queue          *queue.Queue
	SessionHandler *session.Handler

	Config config.GameConfig
}

func New(config config.GameConfig) *Game {
	// idk y 1024
	connPipe := make(chan conn.Connection, 1024)
	seshPipe := make(chan conn.Pair)

	server := server.New(connPipe)
	queue := queue.New(connPipe, seshPipe)
	sHandler := session.NewHandler(seshPipe, config)

	go queue.Accept()
	go queue.Run()
	go sHandler.Run()

	return &Game{
		Server:         server,
		Queue:          queue,
		SessionHandler: sHandler,
	}
}
