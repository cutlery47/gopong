package entities

type Platform struct {
	coord    Vector
	width    float64
	height   float64
	velocity float64
}

func (p *Platform) SetCoord(x, y float64) {
	p.coord.X = x
	p.coord.Y = y
}

func (p Platform) Height() float64 {
	return p.height
}

func (p Platform) Width() float64 {
	return p.width
}

func (p Platform) Coord() Vector {
	return p.coord
}

func (p Platform) Velocity() float64 {
	return p.velocity
}

func (p *Platform) Move(offset float64) {
	p.coord.Y += offset
}

func NewPlatform(width float64, height float64, coord Vector) *Platform {
	p := new(Platform)
	p.width = width
	p.height = height
	p.coord = coord
	p.velocity = 10

	return p
}
