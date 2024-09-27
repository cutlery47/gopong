package session

import (
	"github.com/cutlery47/gopong/common/conn"
	"github.com/cutlery47/gopong/server/config"
)

type Handler struct {
	// session storage
	sessions []Session
	// channel for receiving incoming connection pairs
	seshPipe <-chan conn.Pair

	config config.GameConfig
}

func NewHandler(seshPipe <-chan conn.Pair, config config.GameConfig) *Handler {
	return &Handler{
		sessions: []Session{},
		seshPipe: seshPipe,
		config:   config,
	}
}

func (h *Handler) Run() {
	for {
		conns := <-h.seshPipe
		sesh := InitSession(conns.Conn1, conns.Conn2, h.config)
		go sesh.Run()
	}
}
