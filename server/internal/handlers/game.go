package handlers

import (
	"errors"
	"gopong/server/internal/packet"
)

type GameHandler struct {
	cfg packet.ConfigRequestPacket
}

func (h *GameHandler) Handle(left packet.PlayerStatePacket, right packet.PlayerStatePacket) (newLeft, newRight packet.ServerStatePacket, err error) {
	if !h.configIsSet() {
		return newLeft, newRight, errors.New("config hasn't been set")
	}
	return
}

func (h *GameHandler) configIsSet() bool {
	return h.cfg != (packet.ConfigRequestPacket{})
}

func (h *GameHandler) SetConfig(cfg packet.ConfigRequestPacket) error {
	if h.configIsSet() {
		return errors.New("config has already been set")
	}

	h.cfg = cfg
	return nil
}
