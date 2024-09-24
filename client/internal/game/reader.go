package game

import "github.com/hajimehoshi/ebiten/v2"

type InputReader interface {
	Read() KeyboardInputResult
}

type KeyboardInputReader struct {
	upKey   ebiten.Key
	downKey ebiten.Key
}

func (kir KeyboardInputReader) Read() KeyboardInputResult {
	res := KeyboardInputResult{}

	if ebiten.IsKeyPressed(kir.upKey) {
		res.up = true
	}

	if ebiten.IsKeyPressed(kir.downKey) {
		res.down = true
	}

	return res
}

type KeyboardInputResult struct {
	up   bool
	down bool
}
