package game

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	pipe     inConnPipe
}

func NewServer(pipe inConnPipe) *Server {
	upgrader := websocket.Upgrader{}

	return &Server{
		upgrader: upgrader,
		pipe:     pipe,
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
		s.pipe <- connection{
			conn: c,
		}
	}
}
