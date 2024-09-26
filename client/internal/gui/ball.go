package gui

import (
	"image/color"

	"github.com/cutlery47/gopong/client/internal/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	entity *entities.Ball
	Image  *ebiten.Image
}

func (b *Ball) SetPosition(X, Y float64) {
	b.entity.SetPosition(X, Y)
}

func (b Ball) Coordinates() (float64, float64) {
	return b.entity.Coord().AsTuple()
}

func (b Ball) Velocity() float64 {
	return b.entity.Velocity()
}

func (b Ball) YCoord() float64 {
	return b.entity.Coord().Y
}

func (b Ball) XCoord() float64 {
	return b.entity.Coord().X
}

func (b Ball) Size() int {
	return b.entity.Size()
}

func (b *Ball) Move() {
	b.entity.Move()
}

func (b *Ball) OverlapsLeft(plat *Platform) bool {
	xOverlap := plat.XCoord()-float64(plat.Width()) <= b.XCoord() && b.XCoord() <= plat.XCoord()+float64(plat.Width())
	yOverlap := plat.YCoord() <= b.YCoord()+float64(b.Size()) && b.YCoord()+float64(b.Size()) <= plat.YCoord()+float64(plat.Height())

	return xOverlap && yOverlap
}

func (b *Ball) OverlapsRight(plat *Platform) bool {
	xOverlap := plat.XCoord()-float64(b.Size()) <= b.XCoord() && b.XCoord() <= plat.XCoord()+float64(b.Size())
	yOverlap := plat.YCoord() <= b.YCoord()+float64(b.Size()) && b.YCoord()+float64(b.Size()) <= plat.YCoord()+float64(plat.Height())

	return xOverlap && yOverlap
}

func (b *Ball) OverlapsUpper() bool {
	return b.YCoord() <= 0
}

func (b *Ball) OverlapsLower(height int) bool {
	return b.YCoord() >= float64(height-b.Size())
}

func (b *Ball) PlatformCollide() {
	b.entity.VerticalCollide()
	b.entity.SpeedUp(1.025)
}

func (b *Ball) BorderCollide() {
	b.entity.HorizontalCollide()
}

func NewBall(x_coord float64, y_coord float64, size int) *Ball {
	ball := new(Ball)
	ball.entity = entities.NewBall(size, *entities.NewVector(x_coord, y_coord))
	ball.Image = ebiten.NewImage(size, size)
	ball.Image.Fill(color.RGBA{255, 0, 0, 255})

	return ball
}
