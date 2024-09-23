package game

import "gopong/server/internal/pack"

type Session struct {
	player1 connection
	player2 connection
}

func (s Session) handle() {
	s.assignSides()
	for {
		s.handleMatch()
	}
}

func (s Session) assignSides() {
	leftPack := pack.ServerPacket{
		Status: pack.FoundStatus,
		Side:   pack.LeftSide,
	}

	rightPack := pack.ServerPacket{
		Status: pack.FoundStatus,
		Side:   pack.RightSide,
	}

	s.player1.Send(leftPack)
	s.player2.Send(rightPack)

}

func (s Session) handleMatch() {

}
