package core

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	ebitext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Canvas struct {
	state      *State
	leftImage  *ebiten.Image
	rightImage *ebiten.Image
	ballImage  *ebiten.Image
	scoreText  CanvasScoreText
}

func NewCanvas(state *State) *Canvas {
	ebiten.SetWindowSize(int(state.screen.width), int(state.screen.height))

	leftImage := ebiten.NewImage(int(state.left.width), int(state.left.height))
	rightImage := ebiten.NewImage(int(state.right.width), int(state.right.height))
	ballImage := ebiten.NewImage(int(state.ball.size), int(state.ball.size))
	scoreText := InitCanvasScoreText(10, 10, 24, "")

	leftImage.Fill(color.White)
	rightImage.Fill(color.White)
	ballImage.Fill(color.RGBA{0xff, 0, 0, 0xff})

	return &Canvas{
		state:      state,
		leftImage:  leftImage,
		rightImage: rightImage,
		ballImage:  ballImage,
		scoreText:  scoreText,
	}
}

func (c Canvas) Width() float64 {
	return c.state.screen.width
}

func (c Canvas) Height() float64 {
	return c.state.screen.height
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

func (c Canvas) TextPos() (x, y float64) {
	return c.scoreText.pos.x, c.scoreText.pos.y
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

func (c *Canvas) UpdateScoreText() {
	c.scoreText.text = fmt.Sprintf("%v : %v", c.state.score.left, c.state.score.right)
}

type CanvasScoreText struct {
	pos  vector
	size int
	text string
	face *ebitext.GoTextFace
	src  *ebitext.GoTextFaceSource
}

// idk what halfa dis shie does bru
func InitCanvasScoreText(posX, posY float64, size int, text string) CanvasScoreText {
	src, _ := ebitext.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))

	face := &ebitext.GoTextFace{
		Source: src,
		Size:   24,
	}

	return CanvasScoreText{
		pos:  vector{x: posX, y: posY},
		size: size,
		text: text,
		src:  src,
		face: face,
	}
}
