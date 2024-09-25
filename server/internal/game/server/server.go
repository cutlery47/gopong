package server

import (
	"gopong/server/internal/game/conn"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	// channel for sending received connections
	connPipe chan<- conn.Connection
}

func New(pipe chan<- conn.Connection) *Server {
	upgrader := websocket.Upgrader{}

	return &Server{
		upgrader: upgrader,
		connPipe: pipe,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := s.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("ServeHTTP: ", err)
		w.WriteHeader(400)
	} else {
		log.Printf("Received connection from %v", c.RemoteAddr())
		// senging connection over to the queue
		conn := conn.New(c)
		s.connPipe <- conn
	}
}
