package state

import (
	"gopong/server/internal/entities"

	"github.com/cutlery47/gopong/common/protocol"
)

type State struct {
	Left   *entities.Platform
	Right  *entities.Platform
	Ball   *entities.Ball
	Canvas *entities.Canvas
}

func Init() *State {
	canvasWidth := 1000.0
	canvasHeight := 500.0

	ballSize := 10.0
	platformWidth := 20.0
	platformHeight := 100.0

	ballPosition := entities.NewVector(canvasWidth/2-ballSize/2, canvasHeight/2-ballSize/2)
	leftPosition := entities.NewVector(platformWidth, canvasHeight/2-platformHeight/2)
	rightPosition := entities.NewVector(canvasWidth-2*platformWidth, canvasHeight/2-platformHeight/2)

	Canvas := entities.NewCanvas(canvasWidth, canvasHeight)
	Ball := entities.NewBall(ballSize, *ballPosition)
	Left := entities.NewPlatform(platformWidth, platformHeight, *leftPosition)
	Right := entities.NewPlatform(platformWidth, platformHeight, *rightPosition)

	return &State{
		Left:   Left,
		Right:  Right,
		Ball:   Ball,
		Canvas: Canvas,
	}
}

func (s State) LeftCoord() entities.Vector {
	return s.Left.Coord()
}

func (s State) LeftWidth() float64 {
	return s.Left.Width()
}

func (s State) LeftHeight() float64 {
	return s.Left.Height()
}

func (s State) RightCoord() entities.Vector {
	return s.Right.Coord()
}

func (s State) RightWidth() float64 {
	return s.Right.Width()
}

func (s State) RightHeight() float64 {
	return s.Right.Height()
}

func (s State) BallCoord() entities.Vector {
	return s.Ball.Coord()
}

func (s State) BallSize() float64 {
	return s.Ball.Size()
}

func (s State) CanvasWidth() float64 {
	return s.Canvas.Width()
}

func (s State) CanvasHeight() float64 {
	return s.Canvas.Height()
}

func (s *State) Update(leftInput, rightInput protocol.ClientPacket) {
	if s.Ball.OverlapsLeft(s.Left) || s.Ball.OverlapsRight(s.Right) {
		s.Ball.PlatformCollide()
	}
	if s.Ball.OverlapsUpper() || s.Ball.OverlapsLower(s.CanvasHeight()) {
		s.Ball.BorderCollide()
	}

	s.Ball.Move()
	s.Left.SetCoord(leftInput.Position.X, leftInput.Position.Y)
	s.Right.SetCoord(rightInput.Position.X, rightInput.Position.Y)
}
