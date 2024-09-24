package game

import (
	"gopong/client/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Updater interface {
	Update(window *gui.Window) error
}

type LocalUpdater struct {
	leftReader  InputReader
	rightReader InputReader
}

func NewLocalUpdater() *LocalUpdater {
	return &LocalUpdater{
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

func (lu *LocalUpdater) Update(window *gui.Window) error {
	leftInput := lu.leftReader.Read()
	rightINput := lu.rightReader.Read()

	lu.detectAndHandleCollision(window)

	lu.movePlatform(leftInput, window.Left, window.Height())
	lu.movePlatform(rightINput, window.Right, window.Height())
	lu.moveBall(window.Ball)

	return nil
}

func (lu *LocalUpdater) detectAndHandleCollision(window *gui.Window) {
	if window.Ball.OverlapsLeft(window.Left) || window.Ball.OverlapsRight(window.Right) {
		window.Ball.PlatformCollide()
	}
	if window.Ball.OverlapsUpper() || window.Ball.OverlapsLower(window.Height()) {
		window.Ball.BorderCollide()
	}
}

func (lu *LocalUpdater) movePlatform(res KeyboardInputResult, plat *gui.Platform, height int) {
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

func (u LocalUpdater) moveBall(ball *gui.Ball) {
	ball.Move()
}
