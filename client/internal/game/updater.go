package game

import (
	"gopong/client/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Updater interface {
	Update(window *gui.Window) error
}

type localUpdater struct {
	leftReader  InputReader
	rightReader InputReader
}

func NewLocalUpdater() *localUpdater {
	return &localUpdater{
		leftReader: KeyboardInputReader{
			upKey:   ebiten.KeyW,
			downKey: ebiten.KeyS,
		},
		rightReader: KeyboardInputReader{
			upKey:   ebiten.KeyArrowUp,
			downKey: ebiten.KeyArrowDown,
		},
	}
}

func (lu *localUpdater) Update(window *gui.Window) error {
	leftInput := lu.leftReader.Read()
	rightINput := lu.rightReader.Read()

	lu.detectAndHandleCollision(window)

	lu.movePlatform(leftInput, window.Left, window.Height())
	lu.movePlatform(rightINput, window.Right, window.Height())
	lu.moveBall(window.Ball)

	return nil
}

func (lu *localUpdater) detectAndHandleCollision(window *gui.Window) {
	if window.Ball.OverlapsLeft(window.Left) || window.Ball.OverlapsRight(window.Right) {
		window.Ball.PlatformCollide()
	}
	if window.Ball.OverlapsUpper() || window.Ball.OverlapsLower(window.Height()) {
		window.Ball.BorderCollide()
	}
}

func (lu *localUpdater) movePlatform(res KeyboardInputResult, plat *gui.Platform, height int) {
	var offset float64 = 0

	if res.up {
		offset -= plat.Velocity()
	}
	if res.down {
		offset += plat.Velocity()
	}

	new_coord := plat.YCoord() + offset

	if new_coord >= 0 && new_coord <= float64(height-plat.Height()) {
		plat.Move(offset)
	}
}

func (u localUpdater) moveBall(ball *gui.Ball) {
	ball.Move()
}

// type multiplayerUpdater struct {
// 	statePipe <-chan pack.ServerPacket
// 	inputPipe <-chan KeyboardInputResult
// }

// func NewMultiplayerUpdater(side string) *multiplayerUpdater {
// 	return &multiplayerUpdater{}
// }

// func (mu *multiplayerUpdater) Update(window *gui.Window) error {
// 	input := KeyboardInputResult{}

// 	select {
// 	case data := <-mu.inputPipe:
// 		input = data
// 	default:

// 	}

// 	clientState := pack.ClientPacket{

// 	}

// 	return nil
// }
