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
	state     *state.State
	tickSleep time.Duration

	config config.GameConfig
}

func InitSession(c1, c2 conn.Connection, config config.GameConfig) Session {
	left := initPlayer(c1)
	right := initPlayer(c2)

	tickSleep := time.Duration(1.0/float64(config.GameServerConfig.Tickrate)*1000) * time.Millisecond

	session := Session{
		left:      left,
		right:     right,
		state:     &state.State{},
		config:    config,
		tickSleep: tickSleep,
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
	go s.left.prepare(s.config, "left")
	go s.right.prepare(s.config, "right")

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
	s.state = state.Init(s.config)
	for exit := false; !exit; {
		s.sendCurrentState()

		leftInput := <-s.left.inputCh
		rightInput := <-s.right.inputCh

		exit = s.state.Update(leftInput, rightInput)

		time.Sleep(s.tickSleep)
	}

}

func (s *Session) sendCurrentState() {
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

func (p player) prepare(config config.GameConfig, side protocol.PlayerSide) {
	p.conn.Send(protocol.ServerPacket{Status: protocol.FoundStatus})
	if err := p.conn.ReadACK(); err != nil {
		return
	}

	configPack := protocol.GameConfig{
		Side:           side,
		CanvasWidth:    float64(config.CanvasConfig.Width),
		CanvasHeight:   float64(config.CanvasConfig.Height),
		BallSize:       float64(config.BallConfig.Size),
		PlatformWidth:  float64(config.PlatformConfig.Width),
		PlatformHeight: float64(config.PlatformConfig.Height),
	}

	p.conn.Send(configPack)
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
