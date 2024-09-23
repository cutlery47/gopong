package game

import (
	"gopong/server/internal/pack"
	"log"

	"github.com/gorilla/websocket"
)

// incoming connection
type connection struct {
	conn *websocket.Conn
}

func (c connection) Send(pack pack.ServerPacket) (err error) {
	err = c.conn.WriteJSON(pack)
	if err != nil {
		log.Println("connection.Send():", err)
		return err
	}
	return err
}

func (c connection) Read(pack *pack.ClientPacket) (err error) {
	err = c.conn.ReadJSON(pack)
	if err != nil {
		log.Println("connection.Read():", err)
		return err
	}
	return err
}

func (c connection) SendStatus(status pack.PlayerStatus) {
	pack := pack.ServerPacket{
		Status: status,
	}
	c.Send(pack)
}

// pipes for passing connections back & forth between Queue and Server
type inConnPipe chan<- connection
type outConnPipe <-chan connection

// the game itself
type Game struct {
	Server *Server
	Queue  *Queue
}

func New() *Game {
	// idk why 1024
	connPipe := make(chan connection, 1024)

	server := NewServer(connPipe)
	queue := NewQueue(connPipe)

	go queue.Accept()
	go queue.Run()

	return &Game{
		Server: server,
		Queue:  queue,
	}
}
