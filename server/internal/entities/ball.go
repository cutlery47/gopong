package entities

import (
	"math"
	"math/rand"
	"time"
)

type Ball struct {
	coord Vector
	size  float64
	movec Vector
}

func (b Ball) Coord() Vector {
	return b.coord
}

func (b Ball) Size() float64 {
	return b.size
}

func (b Ball) Velocity() float64 {
	dx := b.movec.X - b.coord.X
	dy := b.movec.Y - b.coord.Y

	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func (b *Ball) Move() {
	b.coord.Add(b.movec)
}

func (b *Ball) SpeedUp(mult float64) {
	b.movec.Mult(mult)
}

func (b *Ball) VerticalCollide() {
	b.movec.X = -b.movec.X
}

func (b *Ball) HorizontalCollide() {
	b.movec.Y = -b.movec.Y
}

func (b *Ball) OverlapsLeft(plat *Platform) bool {
	xOverlap := plat.coord.X-float64(plat.Width()) <= b.coord.X && b.coord.X <= plat.coord.X+float64(plat.Width())
	yOverlap := plat.coord.Y <= b.coord.Y+float64(b.Size()) && b.coord.Y+float64(b.Size()) <= plat.coord.Y+float64(plat.Height())

	return xOverlap && yOverlap
}

func (b *Ball) OverlapsRight(plat *Platform) bool {
	xOverlap := plat.coord.X-float64(b.Size()) <= b.coord.X && b.coord.X <= plat.coord.X+float64(b.Size())
	yOverlap := plat.coord.Y <= b.coord.Y+float64(b.Size()) && b.coord.Y+float64(b.Size()) <= plat.coord.Y+float64(plat.Height())

	return xOverlap && yOverlap
}

func (b *Ball) OverlapsUpper() bool {
	return b.coord.Y <= 0
}

func (b *Ball) OverlapsLower(height float64) bool {
	return b.coord.Y >= height-b.Size()
}

func (b *Ball) PlatformCollide() {
	b.VerticalCollide()
	b.SpeedUp(1.025)
}

func (b *Ball) BorderCollide() {
	b.HorizontalCollide()
}

func NewBall(size float64, coord Vector) *Ball {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	b := new(Ball)
	b.size = size
	b.coord = coord
	b.movec = Vector{}

	// randomizing initial ball direction
	if directX := r.Intn(2); directX == 0 {
		b.movec.X = -5.0
	} else {
		b.movec.X = 5.0
	}

	if directY := r.Intn(2); directY == 0 {
		b.movec.Y = -1.5
	} else {
		b.movec.Y = 1.5
	}

	return b
}
