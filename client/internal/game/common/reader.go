package common

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type KeyboardInputReader struct {
	UpKey   ebiten.Key
	DownKey ebiten.Key
}

var LeftKeyboardInputReader = KeyboardInputReader{
	UpKey:   ebiten.KeyW,
	DownKey: ebiten.KeyS,
}

var RightKeyboardInputReader = KeyboardInputReader{
	UpKey:   ebiten.KeyArrowUp,
	DownKey: ebiten.KeyArrowDown,
}

func (kir KeyboardInputReader) Read() KeyboardInputResult {
	res := KeyboardInputResult{}

	if ebiten.IsKeyPressed(kir.UpKey) {
		res.Up = true
	}

	if ebiten.IsKeyPressed(kir.DownKey) {
		res.Down = true
	}

	log.Println(res)

	return res
}

type KeyboardInputResult struct {
	Up   bool
	Down bool
}
