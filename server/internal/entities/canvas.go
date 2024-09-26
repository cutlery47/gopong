package entities

type Canvas struct {
	width  float64
	height float64
}

func NewCanvas(width, height float64) *Canvas {
	return &Canvas{
		width:  width,
		height: height,
	}
}

func (c Canvas) Width() float64 {
	return c.width
}

func (c Canvas) Height() float64 {
	return c.height
}
