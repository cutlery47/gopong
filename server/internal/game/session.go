package game

import (
	"gopong/server/internal/pack"
	"log"
	"time"
)

type Session struct {
	left  Player
	right Player
	state *State
}

func NewSession(c1, c2 connection) *Session {
	left := Player{
		conn:  c1,
		input: make(chan pack.ClientPacket, 1024),
	}

	right := Player{
		conn:  c2,
		input: make(chan pack.ClientPacket, 1024),
	}

	state := InitState()

	return &Session{
		left:  left,
		right: right,
		state: state,
	}
}

func (s Session) handle() {
	s.assignSides()
	s.handleMatch()
}

func (s Session) assignSides() {
	state := pack.ServerState{
		LeftPosition:  pack.Vector{X: 50, Y: 100},
		RightPosition: pack.Vector{X: 100, Y: 100},
		BallPosition:  pack.Vector{X: 0, Y: 0},
	}

	leftPack := pack.ServerPacket{
		Status: pack.FoundStatus,
		Side:   pack.LeftSide,
		State:  state,
	}

	rightPack := pack.ServerPacket{
		Status: pack.FoundStatus,
		Side:   pack.RightSide,
		State:  state,
	}

	s.left.conn.Send(leftPack)
	s.right.conn.Send(rightPack)

	go s.left.Listen()
	go s.right.Listen()

}

func (s Session) handleMatch() {
	for {
		leftInput := pack.ClientPacket{}
		rightInput := pack.ClientPacket{}

		select {
		case data := <-s.left.input:
			leftInput = data
			log.Println("received left")
		case data := <-s.right.input:
			rightInput = data
			log.Println("received right")
		default:
		}

		s.state.Update(leftInput, rightInput)
		s.SendUpdates()
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Session) SendUpdates() {
	statePack := pack.ServerState{
		LeftPosition:  pack.Vector(s.state.left.Coord()),
		RightPosition: pack.Vector(s.state.right.Coord()),
		BallPosition:  pack.Vector(s.state.ball.Coord()),
	}

	serverPack := pack.ServerPacket{
		Status: pack.PlayingStatus,
		State:  statePack,
		Side:   "xyu",
	}

	s.left.conn.Send(serverPack)
	s.right.conn.Send(serverPack)
}

type Player struct {
	conn  connection
	input chan pack.ClientPacket
}

func (p Player) Listen() {
	for {
		data := pack.ClientPacket{}
		if err := p.conn.Read(&data); err != nil {
			log.Println(err)
		}
		p.input <- data
	}
}
