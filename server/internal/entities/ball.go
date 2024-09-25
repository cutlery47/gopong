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
