package gui

import (
	"gopong/internal/entities"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	entity *entities.Platform
	Image  *ebiten.Image
}

func (p Platform) Coordinates() (float64, float64) {
	return p.entity.Coord().AsTuple()
}

func (p Platform) YCoord() float64 {
	return p.entity.Coord().Y
}

func (p Platform) XCoord() float64 {
	return p.entity.Coord().X
}

func (p Platform) Height() int {
	return p.entity.Height()
}

func (p Platform) Width() int {
	return p.entity.Width()
}

func (p Platform) Velocity() float64 {
	return p.entity.Velocity()
}

func (p *Platform) Move(offset float64) {
	p.entity.Move(offset)
}

func NewPlatform(x_coord float64, y_coord float64, width int, height int) *Platform {
	platform := new(Platform)
	platform.entity = entities.NewPlatform(width, height, *entities.NewVector(x_coord, y_coord))
	platform.Image = ebiten.NewImage(width, height)
	platform.Image.Fill(color.White)

	return platform
}
