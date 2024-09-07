package entities

type Vector struct {
	X float64
	Y float64
}

func (v Vector) AsTuple() (float64, float64) {
	return v.X, v.Y
}

func (v1 *Vector) Add(v2 Vector) {
	v1.X += v2.X
	v1.Y += v2.Y
}

func (v *Vector) Mult(mult float64) {
	v.X *= mult
	v.Y *= mult
}

func NewVector(x float64, y float64) *Vector {
	v := new(Vector)
	v.X = x
	v.Y = y

	return v
}
