package handlers

import (
	"fmt"
	"gopong/server/internal/packet"
	"log"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/websocket"
)

type ConfigHandler struct{}

func (h *ConfigHandler) MatchConfigurations(player1 *websocket.Conn, player2 *websocket.Conn) (cfg packet.ConfigRequestPacket, err error) {
	leftCfg, err := h.recvPlayerConfig(player1)
	if err != nil {
		return packet.ConfigRequestPacket{}, fmt.Errorf("recvIncomingConfig: %v", err)
	}

	rightCfg, err := h.recvPlayerConfig(player2)
	if err != nil {
		return packet.ConfigRequestPacket{}, fmt.Errorf("recvIncomingConfig: %v", err)
	}

	// doesn't matter which one of cofigs is returned as long as they are equal
	return leftCfg, h.sendConfigResponse(player1, player2, cmp.Equal(leftCfg, rightCfg))
}

func (h *ConfigHandler) recvPlayerConfig(player *websocket.Conn) (cfg packet.ConfigRequestPacket, err error) {
	if err = player.ReadJSON(&cfg); err != nil {
		return packet.ConfigRequestPacket{}, fmt.Errorf("recvPlayerConfig: %v", err)
	}

	return cfg, nil
}

func (h *ConfigHandler) sendConfigResponse(player1 *websocket.Conn, player2 *websocket.Conn, okay bool) error {
	res := packet.ConfigResponsePacket{}

	if okay {
		res.Response = packet.ConfigAccept
		res.Left = player1.RemoteAddr().String()
		res.Right = player2.RemoteAddr().String()
	} else {
		log.Printf("players %v and %v configurations do not match\n", player1.RemoteAddr(), player2.RemoteAddr())
		res.Response = packet.ConfigDecline
	}

	if err := player1.WriteJSON(res); err != nil {
		return fmt.Errorf("sendConfigResponse: %v", err)
	}

	if err := player2.WriteJSON(res); err != nil {
		return fmt.Errorf("sendConfigResponse: %v", err)
	}

	return nil
}
