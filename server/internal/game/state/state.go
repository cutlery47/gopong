package state

import (
	"gopong/server/internal/entities"

	"github.com/cutlery47/gopong/common/protocol"
)

type State struct {
	left   *entities.Platform
	right  *entities.Platform
	ball   *entities.Ball
	canvas *entities.Canvas
}

func Init() *State {
	canvasWidth := 1000.0
	canvasHeight := 500.0

	ballSize := 10.0
	platformWidth := 50.0
	platformHeight := 100.0

	ballPosition := entities.NewVector(canvasWidth/2-ballSize/2, canvasHeight/2-ballSize/2)
	leftPosition := entities.NewVector(platformWidth, canvasHeight/2-platformHeight/2)
	rightPosition := entities.NewVector(canvasWidth-2*platformWidth, canvasHeight/2-platformWidth/2)

	canvas := entities.NewCanvas(canvasWidth, canvasWidth)
	ball := entities.NewBall(ballSize, *ballPosition)
	left := entities.NewPlatform(platformWidth, platformHeight, *leftPosition)
	right := entities.NewPlatform(platformWidth, platformHeight, *rightPosition)

	return &State{
		left:   left,
		right:  right,
		ball:   ball,
		canvas: canvas,
	}
}

func (s State) LeftCoord() entities.Vector {
	return s.left.Coord()
}

func (s State) RightCoord() entities.Vector {
	return s.right.Coord()
}

func (s State) BallCoord() entities.Vector {
	return s.ball.Coord()
}

func (s *State) Update(leftInput, rightInput protocol.ClientPacket) {
	s.left.Move(3)
}
