package gui

import (
	"gopong/client/internal/pack"
)

type Window struct {
	height int
	width  int

	Left  *Platform
	Right *Platform
	Ball  *Ball
}

func (w *Window) Update(state pack.ServerPacket) {
	w.Left.SetPosition(state.State.LeftPosition.X, state.State.LeftPosition.Y)
	w.Right.SetPosition(state.State.RightPosition.X, state.State.RightPosition.Y)
	w.Ball.SetPosition(state.State.BallPosition.X, state.State.BallPosition.Y)
}

func (w *Window) Resolution() (int, int) {
	return int(w.width), int(w.height)
}

func (w *Window) Height() int {
	return w.height
}

func (w *Window) Width() int {
	return w.width
}

func NewWindow(width, height int) *Window {
	w := new(Window)

	w.height = height
	w.width = width

	ball_size := width / 50
	plat_width := width / 50
	plat_height := height / 5

	// placing ball in the center
	w.Ball = NewBall(float64(width)/2, float64(height)/2, ball_size)

	// placing both platform on the opposite sides of the screen
	w.Left = NewPlatform(float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)
	w.Right = NewPlatform(float64(width)-2*float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)

	return w
}
