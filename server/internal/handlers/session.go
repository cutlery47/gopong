package handlers

import (
	"fmt"
	"log"
	"time"

	"gopong/server/internal/packet"

	"github.com/gorilla/websocket"
)

type SessionHandler struct {
	gameHandler *GameHandler
}

type Session struct {
	player1 *websocket.Conn
	player2 *websocket.Conn
}

func (h *SessionHandler) Handle(sesh Session, cfg packet.ConfigRequestPacket) {
	log.Printf("Started handling session %v", sesh)

	for {
		left, right, err := h.recvIncomingData(sesh)
		if err != nil {
			log.Println(fmt.Errorf("SessionHandler.Handle: %v", err))
		}

		newLeft, newRight, err := h.handleIncomingData(left, right)
		if err != nil {
			log.Println(fmt.Errorf("SessionHandler.Handle: %v", err))
		}

		err = h.sendUpdatedData(newLeft, newRight)
		if err != nil {
			log.Println(fmt.Errorf("SessionHandler.Handle: %v", err))
		}

		time.Sleep(time.Millisecond * 8)
	}
}

func (h *SessionHandler) recvIncomingData(sesh Session) (left packet.PlayerStatePacket, right packet.PlayerStatePacket, err error) {
	left, err = h.recvPlayerData(sesh.player1)
	if err != nil {
		return packet.PlayerStatePacket{}, packet.PlayerStatePacket{}, fmt.Errorf("recvIncomingData: %v", err)
	}

	right, err = h.recvPlayerData(sesh.player2)
	if err != nil {
		return packet.PlayerStatePacket{}, packet.PlayerStatePacket{}, fmt.Errorf("recvIncomingData: %v", err)
	}

	return left, right, nil
}

func (h *SessionHandler) recvPlayerData(player *websocket.Conn) (data packet.PlayerStatePacket, err error) {
	// if err = player.SetReadDeadline(time.Now().Add(time.Second * 10)); err != nil {
	// 	return data, fmt.Errorf("recvPlayerData: %v", err)
	// }

	if err = player.ReadJSON(&data); err != nil {
		return data, fmt.Errorf("recvPlayerData: %v", err)
	}

	return data, nil
}

func (h *SessionHandler) handleIncomingData(left packet.PlayerStatePacket, right packet.PlayerStatePacket) (newLeft, newRight packet.ServerStatePacket, err error) {
	return h.gameHandler.Handle(left, right)
}

func (h *SessionHandler) sendUpdatedData(left packet.ServerStatePacket, right packet.ServerStatePacket) error {
	return nil
}

func newSessionHandler(gameHandler *GameHandler) *SessionHandler {
	return &SessionHandler{
		gameHandler: gameHandler,
	}
}
