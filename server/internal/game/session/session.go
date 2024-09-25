package session

import (
	"gopong/server/internal/game/conn"
	"gopong/server/internal/game/state"
	"log"
	"time"

	"github.com/cutlery47/gopong/common/protocol"
)

type Session struct {
	left  player
	right player
	state *state.State
}

func Init(c1, c2 conn.Connection) {
	state := state.Init()
	left, right := initPlayers(c1, c2)

	session := Session{
		left:  left,
		right: right,
		state: state,
	}

	session.handle()
}

func (s Session) handle() {
	s.prepareMatch()
	s.handleMatch()
}

func (s Session) prepareMatch() {
	go s.left.Listen()
	go s.right.Listen()
}

func (s Session) handleMatch() {
	for {
		leftInput := protocol.ClientPacket{}
		rightInput := protocol.ClientPacket{}

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
		s.sendUpdatedState()

		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Session) sendUpdatedState() {

	statePack := protocol.ServerState{
		LeftPosition:  protocol.Vector(s.state.LeftCoord()),
		RightPosition: protocol.Vector(s.state.RightCoord()),
		BallPosition:  protocol.Vector(s.state.BallCoord()),
	}

	serverPack := protocol.ServerPacket{
		Status: protocol.PlayingStatus,
		State:  statePack,
	}

	s.left.conn.Send(serverPack)
	s.right.conn.Send(serverPack)
}

type player struct {
	conn  conn.Connection
	input chan protocol.ClientPacket
}

func initPlayers(leftConn, rightConn conn.Connection) (left, right player) {
	left = player{
		conn:  leftConn,
		input: make(chan protocol.ClientPacket),
	}

	right = player{
		conn:  rightConn,
		input: make(chan protocol.ClientPacket),
	}

	return left, right
}

func (p player) Listen() {
	for {
		data := protocol.ClientPacket{}
		if err := p.conn.Read(&data); err != nil {
			log.Println(err)
		}
		p.input <- data
	}
}
