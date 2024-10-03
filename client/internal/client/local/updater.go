package local

// import (
// 	"github.com/cutlery47/gopong/client/internal/game/common"
// 	"github.com/cutlery47/gopong/client/internal/gui"

// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type Updater struct {
// 	leftReader  common.KeyboardInputReader
// 	rightReader common.KeyboardInputReader
// }

// func NewLocalUpdater() *Updater {
// 	return &Updater{
// 		leftReader: common.KeyboardInputReader{
// 			UpKey:   ebiten.KeyW,
// 			DownKey: ebiten.KeyS,
// 		},
// 		rightReader: common.KeyboardInputReader{
// 			UpKey:   ebiten.KeyArrowUp,
// 			DownKey: ebiten.KeyArrowDown,
// 		},
// 	}
// }

// func (u *Updater) Update(canvas *gui.Canvas) error {
// 	leftInput := u.leftReader.Read()
// 	rightINput := u.rightReader.Read()

// 	u.detectAndHandleCollision(canvas)

// 	u.movePlatform(leftInput, canvas.Left(), canvas.Height())
// 	u.movePlatform(rightINput, canvas.Right(), canvas.Height())
// 	u.moveBall(canvas.Ball())

// 	return nil
// }

// func (u *Updater) detectAndHandleCollision(canvas *gui.Canvas) {
// 	if canvas.Ball().OverlapsLeft(canvas.Left()) || canvas.Ball().OverlapsRight(canvas.Right()) {
// 		canvas.Ball().PlatformCollide()
// 	}
// 	if canvas.Ball().OverlapsUpper() || canvas.Ball().OverlapsLower(canvas.Height()) {
// 		canvas.Ball().BorderCollide()
// 	}
// }

// func (u *Updater) movePlatform(res common.KeyboardInputResult, plat *gui.Platform, height int) {
// 	var offset float64 = 0

// 	if res.Up {
// 		offset -= plat.Velocity()
// 	}
// 	if res.Down {
// 		offset += plat.Velocity()
// 	}

// 	newCoord := plat.YCoord() + offset

// 	if newCoord >= 0 && newCoord <= float64(height-plat.Height()) {
// 		plat.Move(offset)
// 	}
// }

// func (u Updater) moveBall(ball *gui.Ball) {
// 	ball.Move()
// }
