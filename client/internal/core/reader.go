package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyboardGameInputReader struct {
	UpKey   ebiten.Key
	DownKey ebiten.Key
}

type KeyboardIdleInputReader struct {
	RestartKey ebiten.Key
	ExitKey    ebiten.Key
}

var LeftKeyboardInputReader = KeyboardGameInputReader{
	UpKey:   ebiten.KeyW,
	DownKey: ebiten.KeyS,
}

var RightKeyboardInputReader = KeyboardGameInputReader{
	UpKey:   ebiten.KeyArrowUp,
	DownKey: ebiten.KeyArrowDown,
}

var IdleKeyboardInputReader = KeyboardIdleInputReader{
	RestartKey: ebiten.KeyR,
	ExitKey:    ebiten.KeyEscape,
}

func (kir KeyboardGameInputReader) Read() KeyboardGameInputResult {
	res := KeyboardGameInputResult{}

	if ebiten.IsKeyPressed(kir.UpKey) {
		res.Up = true
	}

	if ebiten.IsKeyPressed(kir.DownKey) {
		res.Down = true
	}

	return res
}

func (kir KeyboardIdleInputReader) Read() KeyboardIdleInputResult {
	res := KeyboardIdleInputResult{}

	if ebiten.IsKeyPressed(kir.ExitKey) {
		res.Exit = true
	}

	if ebiten.IsKeyPressed(kir.RestartKey) {
		res.Restart = true
	}

	return res
}

type KeyboardGameInputResult struct {
	Up   bool
	Down bool
}

type KeyboardIdleInputResult struct {
	Restart bool
	Exit    bool
}

type CombinedKeyboardGameInputResult struct {
	Left  KeyboardGameInputResult
	Right KeyboardGameInputResult
}
