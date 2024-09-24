package entities

type Platform struct {
	coord    Vector
	width    int
	height   int
	velocity float64
}

func (p Platform) Height() int {
	return p.height
}

func (p Platform) Width() int {
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

func (p *Platform) SetPosition(coord Vector) {
	p.coord = coord
}

func NewPlatform(width int, height int, coord Vector) *Platform {
	p := new(Platform)
	p.width = width
	p.height = height
	p.coord = coord
	p.velocity = 10

	return p
}
