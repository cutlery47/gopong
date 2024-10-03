package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Canvas struct {
	state      *State
	leftImage  *ebiten.Image
	rightImage *ebiten.Image
	ballImage  *ebiten.Image
}

func NewCanvas(state *State) *Canvas {
	ebiten.SetWindowSize(int(state.screen.width), int(state.screen.height))

	leftImage := ebiten.NewImage(int(state.left.width), int(state.left.height))
	rightImage := ebiten.NewImage(int(state.right.width), int(state.right.height))
	ballImage := ebiten.NewImage(int(state.ball.size), int(state.ball.size))

	leftImage.Fill(color.White)
	rightImage.Fill(color.White)
	ballImage.Fill(color.RGBA{0xff, 0, 0, 0xff})

	return &Canvas{
		state:      state,
		leftImage:  leftImage,
		rightImage: rightImage,
		ballImage:  ballImage,
	}
}

func (c Canvas) LeftPos() (x, y float64) {
	return c.state.left.pos.x, c.state.left.pos.y
}

func (c Canvas) RightPos() (x, y float64) {
	return c.state.right.pos.x, c.state.right.pos.y
}

func (c Canvas) BallPos() (x, y float64) {
	return c.state.ball.pos.x, c.state.ball.pos.y
}

func (c Canvas) LeftImage() *ebiten.Image {
	return c.leftImage
}

func (c Canvas) RightImage() *ebiten.Image {
	return c.rightImage
}

func (c Canvas) BallImage() *ebiten.Image {
	return c.ballImage
}
