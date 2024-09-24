package game

import (
	"gopong/client/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawer interface {
	Draw(window *gui.Window, screen *ebiten.Image)
}

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r Renderer) Draw(window *gui.Window, screen *ebiten.Image) {
	op_1 := &ebiten.DrawImageOptions{}
	op_2 := &ebiten.DrawImageOptions{}
	ball_op := &ebiten.DrawImageOptions{}

	op_1.GeoM.Translate(window.Left.Coordinates())
	op_2.GeoM.Translate(window.Right.Coordinates())
	ball_op.GeoM.Translate(window.Ball.Coordinates())

	screen.DrawImage(window.Left.Image, op_1)
	screen.DrawImage(window.Right.Image, op_2)
	screen.DrawImage(window.Ball.Image, ball_op)
}
