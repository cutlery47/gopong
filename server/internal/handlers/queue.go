package handlers

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type QueueHandler struct {
	sHandler *SessionHandler
	queue    []*websocket.Conn
}

func (handler *QueueHandler) Handle() {
	for {
		log.Println("Current connections in queue:", len(handler.queue))
		for len(handler.queue) >= 2 {
			sesh := Session{
				handler.queue[0],
				handler.queue[1],
			}
			handler.queue = handler.queue[2:]
			handler.sHandler.sessions = append(handler.sHandler.sessions, sesh)
			go handler.sHandler.Handle(sesh)
		}
		time.Sleep(1 * time.Second)
	}
}

func newQueueHandler(sHandler *SessionHandler) *QueueHandler {
	queue := []*websocket.Conn{}

	return &QueueHandler{
		queue:    queue,
		sHandler: sHandler,
	}
}
