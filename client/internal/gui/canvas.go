package gui

import (
	"github.com/cutlery47/gopong/common/protocol"
)

type Canvas struct {
	height int
	width  int

	Left  *Platform
	Right *Platform
	Ball  *Ball
}

func (c *Canvas) Update(state protocol.ServerPacket) {
	c.Left.SetPosition(state.State.LeftPosition.X, state.State.LeftPosition.Y)
	c.Right.SetPosition(state.State.RightPosition.X, state.State.RightPosition.Y)
	c.Ball.SetPosition(state.State.BallPosition.X, state.State.BallPosition.Y)
}

func (c *Canvas) Resolution() (int, int) {
	return int(c.width), int(c.height)
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) Width() int {
	return c.width
}

func NewCanvas(width, height int) *Canvas {
	c := &Canvas{}

	c.height = height
	c.width = width

	ball_size := width / 50
	plat_width := width / 50
	plat_height := height / 5

	// placing ball in the center
	c.Ball = NewBall(float64(width)/2, float64(height)/2, ball_size)

	// placing both platform on the opposite sides of the screen
	c.Left = NewPlatform(float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)
	c.Right = NewPlatform(float64(width)-2*float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)

	return c
}

func NewCanvasFromConfig(config protocol.GameConfig) *Canvas {
	c := &Canvas{}

	c.height = int(config.CanvasHeight)
	c.width = int(config.CanvasWidth)

	c.Ball = NewBall(config.BallPosition.X, config.BallPosition.Y, int(config.BallSize))
	c.Left = NewPlatform(config.LeftPlatformPosition.X, config.LeftPlatformPosition.Y, int(config.LeftWidth), int(config.LeftHeight))
	c.Right = NewPlatform(config.RightPlatformPosition.X, config.RightPlatformPosition.Y, int(config.RightWidth), int(config.RightHeight))

	return c
}
