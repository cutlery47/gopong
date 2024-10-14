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
	scoreText  CanvasText
}

func NewCanvas(state *State) *Canvas {
	ebiten.SetWindowSize(int(state.screen.width), int(state.screen.height))

	leftImage := ebiten.NewImage(int(state.left.width), int(state.left.height))
	rightImage := ebiten.NewImage(int(state.right.width), int(state.right.height))
	ballImage := ebiten.NewImage(int(state.ball.size), int(state.ball.size))
	scoreText := InitCanvasText(10, 10, 24, "")

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

type IdleCanvas struct {
	resultText    CanvasText
	instructText1 CanvasText
	instructText2 CanvasText
}

func (i IdleCanvas) ResPos() (x, y float64) {
	return i.resultText.pos.x, i.resultText.pos.y
}

func (i IdleCanvas) InstPos1() (x, y float64) {
	return i.instructText1.pos.x, i.instructText1.pos.y
}

func (i IdleCanvas) InstPos2() (x, y float64) {
	return i.instructText2.pos.x, i.instructText2.pos.y
}

func NewIdleCanvas() *IdleCanvas {
	resultText := InitCanvasText(0, 0, 30, "")
	instructText1 := InitCanvasText(0, 100, 30, "Press R to Start the Game")
	instructText2 := InitCanvasText(0, 150, 30, "Press ESC to Quit the Game")

	return &IdleCanvas{
		resultText:    resultText,
		instructText1: instructText1,
		instructText2: instructText2,
	}
}

type CanvasText struct {
	pos  vector
	size int
	text string
	face *ebitext.GoTextFace
	src  *ebitext.GoTextFaceSource
}

// idk what halfa dis shie does bru
func InitCanvasText(posX, posY float64, size int, text string) CanvasText {
	src, _ := ebitext.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))

	face := &ebitext.GoTextFace{
		Source: src,
		Size:   float64(size),
	}

	return CanvasText{
		pos:  vector{x: posX, y: posY},
		size: size,
		text: text,
		src:  src,
		face: face,
	}
}
