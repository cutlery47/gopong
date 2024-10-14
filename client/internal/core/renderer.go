package core

import (
	"github.com/hajimehoshi/ebiten/v2"

	ebitext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Renderer struct {
	canvas     *Canvas
	idleCanvas *IdleCanvas
}

func NewRenderer(canvas *Canvas, idleCanvas *IdleCanvas) *Renderer {
	return &Renderer{
		canvas:     canvas,
		idleCanvas: idleCanvas,
	}
}

func (r *Renderer) DrawGame(screen *ebiten.Image) {
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

func (r *Renderer) DrawIdle(screen *ebiten.Image) {
	resOp := &ebitext.DrawOptions{}
	instOp1 := &ebitext.DrawOptions{}
	instOp2 := &ebitext.DrawOptions{}

	resOp.GeoM.Translate(r.idleCanvas.ResPos())
	instOp1.GeoM.Translate(r.idleCanvas.InstPos1())
	instOp2.GeoM.Translate(r.idleCanvas.InstPos2())

	ebitext.Draw(screen, r.idleCanvas.resultText.text, r.idleCanvas.resultText.face, resOp)
	ebitext.Draw(screen, r.idleCanvas.instructText1.text, r.idleCanvas.instructText1.face, instOp1)
	ebitext.Draw(screen, r.idleCanvas.instructText2.text, r.idleCanvas.instructText2.face, instOp2)
}
