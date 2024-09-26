package common

import (
	"github.com/cutlery47/gopong/client/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r Renderer) Draw(canvas *gui.Canvas, screen *ebiten.Image) {
	op_1 := &ebiten.DrawImageOptions{}
	op_2 := &ebiten.DrawImageOptions{}
	ball_op := &ebiten.DrawImageOptions{}

	// locating objects on the canvas
	op_1.GeoM.Translate(canvas.Left().Coordinates())
	op_2.GeoM.Translate(canvas.Right().Coordinates())
	ball_op.GeoM.Translate(canvas.Ball().Coordinates())

	// rendering
	screen.DrawImage(canvas.Left().Image, op_1)
	screen.DrawImage(canvas.Right().Image, op_2)
	screen.DrawImage(canvas.Ball().Image, ball_op)
}
