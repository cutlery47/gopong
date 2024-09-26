package gui

import (
	"github.com/cutlery47/gopong/common/protocol"
)

type Canvas struct {
	height int
	width  int

	left  *Platform
	right *Platform
	ball  *Ball
}

func (c Canvas) Resolution() (int, int) {
	return c.width, c.height
}

func (c Canvas) Height() int {
	return c.height
}

func (c Canvas) Width() int {
	return c.width
}

func (c Canvas) Left() *Platform {
	return c.left
}

func (c Canvas) Right() *Platform {
	return c.right
}

func (c Canvas) Ball() *Ball {
	return c.ball
}

func NewCanvas(width, height int) *Canvas {
	c := &Canvas{}

	c.height = height
	c.width = width

	ball_size := width / 50
	plat_width := width / 50
	plat_height := height / 5

	// placing ball in the center
	c.ball = NewBall(float64(width)/2, float64(height)/2, ball_size)

	// placing both platform on the opposite sides of the screen
	c.left = NewPlatform(float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)
	c.right = NewPlatform(float64(width)-2*float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)

	return c
}

func NewCanvasFromConfig(config protocol.GameConfig) *Canvas {
	c := &Canvas{}

	c.height = int(config.CanvasHeight)
	c.width = int(config.CanvasWidth)

	c.ball = NewBall(config.BallPosition.X, config.BallPosition.Y, int(config.BallSize))
	c.left = NewPlatform(config.LeftPlatformPosition.X, config.LeftPlatformPosition.Y, int(config.LeftWidth), int(config.LeftHeight))
	c.right = NewPlatform(config.RightPlatformPosition.X, config.RightPlatformPosition.Y, int(config.RightWidth), int(config.RightHeight))

	return c
}
