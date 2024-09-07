package gui

type Window struct {
	height int
	width  int

	Platform_1 *Platform
	Platform_2 *Platform
	Ball       *Ball
}

func (w *Window) Resolution() (int, int) {
	return int(w.width), int(w.height)
}

func (w *Window) Height() int {
	return w.height
}

func (w *Window) Width() int {
	return w.width
}

func NewWindow(width, height int) *Window {
	w := new(Window)

	w.height = height
	w.width = width

	ball_size := width / 50
	plat_width := width / 50
	plat_height := height / 5

	// placing ball in the center
	w.Ball = NewBall(float64(width)/2, float64(height)/2, ball_size)

	// placing both platform on the opposite sides of the screen
	w.Platform_1 = NewPlatform(float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)
	w.Platform_2 = NewPlatform(float64(width)-2*float64(plat_width), float64(height)/2-float64(plat_height)/2, plat_width, plat_height)

	return w
}
