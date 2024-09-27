package session

import (
	"log"
	"time"

	"github.com/cutlery47/gopong/server/config"
	"github.com/cutlery47/gopong/server/internal/game/state"

	"github.com/cutlery47/gopong/common/conn"
	"github.com/cutlery47/gopong/common/protocol"
)

type Session struct {
	left      player
	right     player
	currState *state.State
	initState *state.State
}

func InitSession(c1, c2 conn.Connection, config config.GameConfig) Session {
	initState := state.Init(config)
	left := initPlayer(c1)
	right := initPlayer(c2)

	session := Session{
		left:      left,
		right:     right,
		currState: &state.State{},
		initState: initState,
	}

	return session
}

func (s Session) Run() {
	s.prepareMatch()
	for {
		s.handleMatch()
	}

}

func (s Session) prepareMatch() {
	go s.left.prepare(*s.initState, "left")
	go s.right.prepare(*s.initState, "right")

	leftReady := false
	rightReady := false

	for !leftReady || !rightReady {
		select {
		case <-s.left.prepCh:
			leftReady = true
		case <-s.right.prepCh:
			rightReady = true
		}
	}
}

func (s Session) handleMatch() {
	*s.currState = *s.initState
	for {
		s.sendCurrentState()

		leftInput := <-s.left.inputCh
		rightInput := <-s.right.inputCh

		if exit := s.currState.Update(leftInput, rightInput); exit {
			return
		}

		// approx 128 tickrate
		time.Sleep(8 * time.Millisecond)
	}
}

func (s *Session) sendCurrentState() {
	statePack := protocol.ServerState{
		LeftPosition:  protocol.Vector(s.currState.LeftCoord()),
		RightPosition: protocol.Vector(s.currState.RightCoord()),
		BallPosition:  protocol.Vector(s.currState.BallCoord()),
	}

	serverPack := protocol.ServerPacket{
		Status: protocol.PlayingStatus,
		State:  statePack,
	}

	s.left.conn.Send(serverPack)
	s.right.conn.Send(serverPack)
}

type player struct {
	conn    conn.Connection
	inputCh chan protocol.ClientPacket
	prepCh  chan byte
}

func initPlayer(conn conn.Connection) player {
	player := player{
		conn:    conn,
		inputCh: make(chan protocol.ClientPacket),
		prepCh:  make(chan byte),
	}

	return player
}

func (p player) prepare(state state.State, side protocol.PlayerSide) {
	p.conn.Send(protocol.ServerPacket{Status: protocol.FoundStatus})
	if err := p.conn.ReadACK(); err != nil {
		return
	}

	config := protocol.GameConfig{
		Side:                  side,
		CanvasWidth:           state.CanvasWidth(),
		CanvasHeight:          state.CanvasHeight(),
		BallSize:              state.BallSize(),
		LeftWidth:             state.LeftWidth(),
		LeftHeight:            state.LeftHeight(),
		RightHeight:           state.RightHeight(),
		RightWidth:            state.RightWidth(),
		BallPosition:          protocol.Vector(state.BallCoord()),
		LeftPlatformPosition:  protocol.Vector(state.LeftCoord()),
		RightPlatformPosition: protocol.Vector(state.RightCoord()),
	}

	p.conn.Send(config)
	if err := p.conn.ReadACK(); err != nil {
		return
	}

	p.prepCh <- 1

	go p.listen()
}

func (p player) listen() {
	for {
		data := protocol.ClientPacket{}
		if err := p.conn.Read(&data); err != nil {
			log.Println(err)
		}
		p.inputCh <- data
	}
}
