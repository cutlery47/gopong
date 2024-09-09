package handlers

import (
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/gorilla/websocket"
)

type QueueHandler struct {
	configHandler  *ConfigHandler
	sessionHandler *SessionHandler
	queue          []*websocket.Conn
}

func (h *QueueHandler) Handle() {
	for {
		log.Println("Current connections in queue:", len(h.queue))
		h.handleQueue()
		time.Sleep(1 * time.Second)
	}
}

func (h *QueueHandler) handleQueue() {
	// iterating over each connection in the queue in order to create a gaming session
	// in order for gaming session to be created, player configs should be match
	// otherwise, session will be discarded

	i := 0
	for len(h.queue) >= 2 && i < len(h.queue)-1 {
		firstConn := h.queue[i]
		secondConn := h.queue[i+1]

		// check if configurations are the same
		cfg, err := h.configHandler.MatchConfigurations(firstConn, secondConn)
		if err != nil {
			// if not - try the next pair of connections
			log.Println(fmt.Errorf("QueueHandler.Handle: %v", err))
			i++
			continue
		}

		// create a new session and start handling it
		sesh := Session{
			player1: firstConn,
			player2: secondConn,
		}
		go h.sessionHandler.Handle(sesh, cfg)

		// update the queue
		h.queue = slices.Concat(h.queue[:i], h.queue[i+2:])
		i = 0
	}
}

func newQueueHandler(sessionHandler *SessionHandler, configHandler *ConfigHandler) *QueueHandler {
	queue := []*websocket.Conn{}

	return &QueueHandler{
		queue:          queue,
		sessionHandler: sessionHandler,
		configHandler:  configHandler,
	}
}
