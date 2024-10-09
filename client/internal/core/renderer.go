package core

import (
	"github.com/hajimehoshi/ebiten/v2"

	ebitext "github.com/hajimehoshi/ebiten/v2/text/v2"
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
	op1 := &ebiten.DrawImageOptions{}
	op2 := &ebiten.DrawImageOptions{}
	ballOp := &ebiten.DrawImageOptions{}
	msgOp := &ebitext.DrawOptions{}

	// locating objects on the canvas
	op1.GeoM.Translate(r.canvas.LeftPos())
	op2.GeoM.Translate(r.canvas.RightPos())
	ballOp.GeoM.Translate(r.canvas.BallPos())
	msgOp.GeoM.Translate(r.canvas.TextPos())

	// scoreMsg := fmt.Sprintf("%v : %v", r.canvas.state.score.left, r.canvas.state.score.right)

	r.canvas.UpdateScoreText()

	// rendering
	screen.DrawImage(r.canvas.LeftImage(), op1)
	screen.DrawImage(r.canvas.RightImage(), op2)
	screen.DrawImage(r.canvas.BallImage(), ballOp)
	ebitext.Draw(screen, r.canvas.scoreText.text, r.canvas.scoreText.face, msgOp)
}
