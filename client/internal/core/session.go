package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Session struct {
	renderer    *Renderer
	leftReader  KeyboardInputReader
	rightReader KeyboardInputReader

	// channel for sending game updates
	inputChan chan<- CombinedKeyboardInputResult
}

func InitSession(inputChan chan<- CombinedKeyboardInputResult, state *State) Session {
	canvas := NewCanvas(state)
	renderer := NewRenderer(canvas)

	// left player reader
	leftReader := LeftKeyboardInputReader
	// right player reader
	rightReader := RightKeyboardInputReader

	return Session{
		renderer:    renderer,
		leftReader:  leftReader,
		rightReader: rightReader,
		inputChan:   inputChan,
	}
}

func (s Session) Update() error {
	// receiving keyboard input
	leftInput := s.leftReader.Read()
	rightInput := s.rightReader.Read()
	// sending the input
	s.inputChan <- CombinedKeyboardInputResult{Left: leftInput, Right: rightInput}
	return nil
}

func (s Session) Draw(screen *ebiten.Image) {
	s.renderer.Draw(screen)
}

func (s Session) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
