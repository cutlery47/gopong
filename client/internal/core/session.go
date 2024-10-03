package core

import (
	"log"

	"github.com/cutlery47/gopong/client/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type Session struct {
	state       *State
	renderer    *Renderer
	leftReader  KeyboardInputReader
	rightReader KeyboardInputReader

	// channel for receiving game updates
	updateChan <-chan StateUpdate
	// channel for sending game updates
	inputChan chan<- CombinedKeyboardInputResult
}

func InitSession(updateChan <-chan StateUpdate, inputChan chan<- CombinedKeyboardInputResult, cfg config.LocalConfig, fpsLimit int) Session {
	// creating shared entities
	ball := &ball{}
	left := &platform{}
	right := &platform{}

	canvas := NewCanvas(
		cfg.ScreenWidth,
		cfg.ScreenHeight,
		cfg.PlatWidth,
		cfg.PlatHeight,
		cfg.BallSize,
		left,
		right,
		ball)
	renderer := NewRenderer(canvas)
	state := NewState(ball, left, right)
	// left player reader
	leftReader := LeftKeyboardInputReader
	// right player reader
	rightReader := RightKeyboardInputReader

	ebiten.SetTPS(fpsLimit)

	return Session{
		renderer:    renderer,
		state:       state,
		leftReader:  leftReader,
		rightReader: rightReader,
		updateChan:  updateChan,
		inputChan:   inputChan,
	}
}

func (s Session) Update() error {
	// receiving update and updating state
	upd := <-s.updateChan
	s.state.Update(upd)
	log.Printf("left: %+v\n", s.state.left)
	log.Printf("right: %+v\n", s.state.right)
	// receiving keyboard input
	leftInput := s.leftReader.Read()
	rightInput := s.rightReader.Read()
	// sending the input
	s.inputChan <- CombinedKeyboardInputResult{Left: leftInput, Right: rightInput}
	return nil
}

func (s Session) Draw(screen *ebiten.Image) {
	s.renderer.Draw(*s.state, screen)
}

func (s Session) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
