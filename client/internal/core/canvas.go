package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Canvas struct {
	width  int
	height int
	screen *ebiten.Image

	left  canvasPlatform
	right canvasPlatform
	ball  canvasBall
}

func NewCanvas(width, height, platWidth, platHeight, ballSize int, left, right *platform, ball *ball) *Canvas {
	screen := ebiten.NewImage(width, height)
	leftCanvasPlatform := initCanvasPlatform(platWidth, platHeight, left)
	rightCanvasPlatform := initCanvasPlatform(platWidth, platHeight, right)
	canvasBall := initCanvasBall(ballSize, ball)

	return &Canvas{
		width:  width,
		height: height,
		screen: screen,
		ball:   canvasBall,
		left:   leftCanvasPlatform,
		right:  rightCanvasPlatform,
	}
}

func (c Canvas) LeftPos() (x, y float64) {
	return c.left.plat.pos.x, c.left.plat.pos.y
}

func (c Canvas) RightPos() (x, y float64) {
	return c.right.plat.pos.x, c.right.plat.pos.y
}

func (c Canvas) BallPos() (x, y float64) {
	return c.ball.ball.pos.x, c.ball.ball.pos.y
}

func (c Canvas) LeftImage() *ebiten.Image {
	return c.left.image
}

func (c Canvas) RightImage() *ebiten.Image {
	return c.right.image
}

func (c Canvas) BallImage() *ebiten.Image {
	return c.ball.image
}

type canvasPlatform struct {
	width  int
	height int

	image *ebiten.Image
	plat  *platform
}

func initCanvasPlatform(width, height int, plat *platform) canvasPlatform {
	image := ebiten.NewImage(width, height)
	image.Fill(color.White)

	return canvasPlatform{
		width:  width,
		height: height,
		image:  image,
		plat:   plat,
	}
}

type canvasBall struct {
	size int

	image *ebiten.Image
	ball  *ball
}

func initCanvasBall(size int, ball *ball) canvasBall {
	image := ebiten.NewImage(size, size)
	image.Fill(color.RGBA{0xff, 0, 0, 0xff})

	return canvasBall{
		size:  size,
		image: image,
		ball:  ball,
	}
}
