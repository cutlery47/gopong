package state

import (
	"github.com/cutlery47/gopong/server/config"
	"github.com/cutlery47/gopong/server/internal/entities"

	"github.com/cutlery47/gopong/common/protocol"
)

type State struct {
	Left   entities.Platform
	Right  entities.Platform
	Ball   entities.Ball
	Canvas entities.Canvas
}

func Init(config config.GameConfig) *State {
	canvasWidth := float64(config.CanvasConfig.Width)
	canvasHeight := float64(config.CanvasConfig.Height)

	ballSize := float64(config.BallConfig.Size)
	platformWidth := float64(config.PlatformConfig.Width)
	platformHeight := float64(config.PlatformConfig.Height)
	platformSpeed := float64(config.PlatformConfig.Speed)

	ballPosition := entities.NewVector(canvasWidth/2-ballSize/2, canvasHeight/2-ballSize/2)
	leftPosition := entities.NewVector(platformWidth, canvasHeight/2-platformHeight/2)
	rightPosition := entities.NewVector(canvasWidth-2*platformWidth, canvasHeight/2-platformHeight/2)

	Canvas := entities.NewCanvas(canvasWidth, canvasHeight)
	Ball := entities.NewBall(ballSize, *ballPosition)
	Left := entities.NewPlatform(platformWidth, platformHeight, *leftPosition, platformSpeed)
	Right := entities.NewPlatform(platformWidth, platformHeight, *rightPosition, platformSpeed)

	return &State{
		Left:   *Left,
		Right:  *Right,
		Ball:   *Ball,
		Canvas: *Canvas,
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

func (s *State) Update(leftInput, rightInput protocol.ClientPacket) bool {
	if s.Ball.OverlapsLeft(s.Left) || s.Ball.OverlapsRight(s.Right) {
		s.Ball.PlatformCollide()
	}
	if s.Ball.OverlapsUpper() || s.Ball.OverlapsLower(s.CanvasHeight()) {
		s.Ball.BorderCollide()
	}

	leftOffest := leftInput.Position.Y - s.LeftCoord().Y
	rightOffset := rightInput.Position.Y - s.RightCoord().Y

	s.Ball.Move()
	if s.BallCoord().X <= 0 || s.BallCoord().X >= s.CanvasWidth() {
		return true
	}

	if 0 <= s.Left.Coord().Y+leftOffest && s.LeftCoord().Y+leftOffest+s.Left.Height() <= s.CanvasHeight() {
		s.Left.Move(leftOffest)
	}
	if 0 <= s.Right.Coord().Y+rightOffset && s.RightCoord().Y+rightOffset+s.Right.Height() <= s.CanvasHeight() {
		s.Right.Move(rightOffset)
	}

	return false
}
