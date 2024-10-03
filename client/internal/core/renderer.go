package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	canvas *Canvas
}

func NewRenderer(canvas *Canvas) *Renderer {
	return &Renderer{
		canvas: canvas,
	}
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	// log.Printf("%+v\n", r.canvas.state)
	op_1 := &ebiten.DrawImageOptions{}
	op_2 := &ebiten.DrawImageOptions{}
	ball_op := &ebiten.DrawImageOptions{}

	// locating objects on the canvas
	op_1.GeoM.Translate(r.canvas.LeftPos())
	op_2.GeoM.Translate(r.canvas.RightPos())
	ball_op.GeoM.Translate(r.canvas.BallPos())

	// rendering
	screen.DrawImage(r.canvas.LeftImage(), op_1)
	screen.DrawImage(r.canvas.RightImage(), op_2)
	screen.DrawImage(r.canvas.BallImage(), ball_op)
}
