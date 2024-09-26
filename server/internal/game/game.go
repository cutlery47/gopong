package game

import (
	"gopong/server/internal/game/queue"
	"gopong/server/internal/game/server"

	"github.com/cutlery47/gopong/common/conn"
)

// the game itself
type Game struct {
	Server *server.Server
	Queue  *queue.Queue
}

func New() *Game {
	// idk y 1024
	connPipe := make(chan conn.Connection, 1024)

	server := server.New(connPipe)
	queue := queue.New(connPipe)

	go queue.Accept()
	go queue.Run()

	return &Game{
		Server: server,
		Queue:  queue,
	}
}
