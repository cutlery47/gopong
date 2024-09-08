package game

import (
	"gopong/client/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawer interface {
	Draw(*gui.Platform, *gui.Platform, *gui.Ball, *ebiten.Image)
}

type GameDrawer struct{}

func (d GameDrawer) Draw(left *gui.Platform, right *gui.Platform, ball *gui.Ball, screen *ebiten.Image) {
	op_1 := &ebiten.DrawImageOptions{}
	op_2 := &ebiten.DrawImageOptions{}
	ball_op := &ebiten.DrawImageOptions{}

	op_1.GeoM.Translate(left.Coordinates())
	op_2.GeoM.Translate(right.Coordinates())
	ball_op.GeoM.Translate(ball.Coordinates())

	screen.DrawImage(left.Image, op_1)
	screen.DrawImage(right.Image, op_2)
	screen.DrawImage(ball.Image, ball_op)
}
