package core

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var SessionGameMode string = "inGame"
var SessionIdleMode string = "inIdle"
var SessionQueueMode string = "inQueue"

type Session struct {
	renderer    *Renderer
	leftReader  KeyboardGameInputReader
	rightReader KeyboardGameInputReader
	idleReader  KeyboardIdleInputReader
	mode        string

	// channel for sending game updates
	inputChan chan<- CombinedKeyboardGameInputResult
	// channel for closing a game
	exitChan <-chan byte
	// channel for signaling that game is about to close
	finishChan chan<- byte
	// channel for signaling that game is about to restart
	startChan chan<- byte
	// channel for setting session mode to idle
	idleChan <-chan byte
}

func InitSession(inputChan chan<- CombinedKeyboardGameInputResult, exitChan <-chan byte, finishChan, startChan, idleChan chan<- byte, state *State) Session {
	canvas := NewCanvas(state)
	renderer := NewRenderer(canvas)

	// left player reader
	leftReader := LeftKeyboardInputReader
	// right player reader
	rightReader := RightKeyboardInputReader
	// idle mode reader
	idleReader := IdleKeyboardInputReader

	return Session{
		renderer:    renderer,
		leftReader:  leftReader,
		rightReader: rightReader,
		idleReader:  idleReader,
		inputChan:   inputChan,
		exitChan:    exitChan,
		finishChan:  finishChan,
		startChan:   startChan,
		mode:        SessionIdleMode,
	}
}

// reading and handling user input
func (s Session) Update() error {
	if exit := s.listenForExit(); exit != nil {
		return exit
	}

	switch s.mode {
	case SessionIdleMode:
		return s.UpdateIdle()
	case SessionGameMode:
		return s.UpdateGame()
	}

	panic("update: unreachable")
}

// rendering
func (s Session) Draw(screen *ebiten.Image) {
	switch s.mode {
	case SessionIdleMode:
		s.renderer.DrawIdle(screen)
	case SessionGameMode:
		s.renderer.DrawGame(screen)
	}
}

// reading and handling user input while game is active
func (s Session) UpdateGame() error {
	log.Println("xwsdf")
	leftInput := s.leftReader.Read()
	rightInput := s.rightReader.Read()
	// sending the input
	s.inputChan <- CombinedKeyboardGameInputResult{Left: leftInput, Right: rightInput}

	return nil
}

// reading and hanlding user input while game is idle
func (s *Session) UpdateIdle() error {
	idleInput := s.idleReader.Read()
	if idleInput.Exit {
		s.finishChan <- 1
	} else if idleInput.Restart {
		s.startChan <- 1
		s.mode = SessionGameMode
	}

	return nil
}

func (s Session) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (s Session) listenForExit() error {
	select {
	case <-s.exitChan:
		return ebiten.Termination
	default:
		return nil
	}
}

func (s *Session) ListenForIdle() {
	for {
		<-s.idleChan
		s.mode = SessionIdleMode
	}
}
