package game

import "log"

type Updater interface {
	Update() error
}

type DefaultUpdater struct{}

func NewDefaultUpdater() Updater {
	return &DefaultUpdater{}
}

func (du *DefaultUpdater) Update() error {
	log.Println("update")
	return nil
}

// type Updater interface {
// 	Update(*gui.Platform, *gui.Platform, *gui.Ball) error
// }

// type LocalUpdater struct {
// 	reader InputReader
// }

// func (u LocalUpdater) Update(cfg *config.Config, left *gui.Platform, right *gui.Platform, ball *gui.Ball) error {
// 	u.detectAndHandleCollision(cfg, left, right, ball)

// 	left_offset := u.calcPlatformOffset(left.Velocity(), u.reader.ReadLeft())
// 	right_offset := u.calcPlatformOffset(right.Velocity(), u.reader.ReadRight())

// 	u.movePlatform(left, left_offset, cfg.Window.Height)
// 	u.movePlatform(right, right_offset, cfg.Window.Height)
// 	u.moveBall(ball)

// 	return nil
// }

// // detects ball-platform collision
// // if detected - returns the platform which has collided with the ball
// // else nil

// func (u LocalUpdater) detectAndHandleCollision(cfg *config.Config, left *gui.Platform, right *gui.Platform, ball *gui.Ball) {
// 	if u.detectPlatformCollision(left, right, ball) {
// 		u.handlePlatformCollision(ball)
// 	} else if u.detectBorderCollision(cfg, ball) {
// 		u.handleBorderCollision(ball)
// 	}
// }

// func (u LocalUpdater) detectPlatformCollision(left *gui.Platform, right *gui.Platform, ball *gui.Ball) bool {
// 	if ball.OverlapsLeft(left) || ball.OverlapsRight(right) {
// 		return true
// 	}
// 	return false
// }

// func (u LocalUpdater) handlePlatformCollision(ball *gui.Ball) {
// 	ball.PlatformCollide()
// }

// func (u LocalUpdater) detectBorderCollision(cfg *config.Config, ball *gui.Ball) bool {
// 	if ball.OverlapsUpper() || ball.OverlapsLower(cfg.Window.Height) {
// 		return true
// 	}
// 	return false
// }

// func (u LocalUpdater) handleBorderCollision(ball *gui.Ball) {
// 	ball.BorderCollide()
// }

// func (u LocalUpdater) calcPlatformOffset(vel float64, res KeyboardInputResult) float64 {
// 	var offset float64
// 	if res.up {
// 		offset -= vel
// 	}
// 	if res.down {
// 		offset += vel
// 	}
// 	return offset
// }

// func (u LocalUpdater) movePlatform(plat *gui.Platform, offset float64, height int) {
// 	new_coord := plat.YCoord() + offset

// 	if new_coord >= 0 && new_coord <= float64(height-plat.Height()) {
// 		plat.Move(offset)
// 	}
// }

// func (u LocalUpdater) moveBall(ball *gui.Ball) {
// 	ball.Move()
// }

// func newLocalUpdater() *LocalUpdater {
// 	updater := &LocalUpdater{
// 		reader: KeyboardInputReader{},
// 	}

// 	return updater
// }
