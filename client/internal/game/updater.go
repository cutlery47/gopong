package game

import (
	"github.com/cutlery47/gopong/client/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Updater interface {
	Update(canvas *gui.Canvas) error
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

func (lu *localUpdater) Update(canvas *gui.Canvas) error {
	leftInput := lu.leftReader.Read()
	rightINput := lu.rightReader.Read()

	lu.detectAndHandleCollision(canvas)

	lu.movePlatform(leftInput, canvas.Left, canvas.Height())
	lu.movePlatform(rightINput, canvas.Right, canvas.Height())
	lu.moveBall(canvas.Ball)

	return nil
}

func (lu *localUpdater) detectAndHandleCollision(canvas *gui.Canvas) {
	if canvas.Ball.OverlapsLeft(canvas.Left) || canvas.Ball.OverlapsRight(canvas.Right) {
		canvas.Ball.PlatformCollide()
	}
	if canvas.Ball.OverlapsUpper() || canvas.Ball.OverlapsLower(canvas.Height()) {
		canvas.Ball.BorderCollide()
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

// func (mu *multiplayerUpdater) Update(canvas *gui.Window) error {
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
