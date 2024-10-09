package core

import (
	"math/rand/v2"

	"github.com/cutlery47/gopong/client/config"
)

type State struct {
	ball   ball
	left   platform
	right  platform
	screen screen
	score  score
}

func NewState() *State {
	return &State{
		ball:   ball{},
		left:   platform{},
		right:  platform{},
		screen: screen{},
	}
}

func StateFromConfig(conf config.StateConfig) State {
	return State{
		ball:   initBall(conf.BallSize, conf.BallInitVelX, conf.BallInitVelY, conf.BallAccelMult),
		left:   platform{width: conf.PlatWidth, height: conf.PlatHeight, vel: conf.PlatVelocity},
		right:  platform{width: conf.PlatWidth, height: conf.PlatHeight, vel: conf.PlatVelocity},
		screen: screen{width: conf.ScreenWidth, height: conf.ScreenHeight},
		score:  score{left: 0, right: 0, max: conf.PointsToWin},
	}
}

type score struct {
	left  int
	right int
	max   int
}

func (s *State) Flush() {
	s.left.pos.x = 0
	s.left.pos.y = s.screen.height/2 - s.left.height/2
	s.right.pos.x = s.screen.width - s.right.width
	s.right.pos.y = s.screen.height/2 - s.left.height/2
	s.ball.pos.x = s.screen.width/2 - s.ball.size/2
	s.ball.pos.y = s.screen.height/2 - s.ball.size/2
}

func (s *State) LeftMoveUp() {
	if s.left.pos.y-s.left.vel > 0 {
		s.left.MoveUp()
	} else {
		s.left.pos.y = 0
	}
}

func (s *State) LeftMoveDown() {
	if s.left.pos.y+s.left.vel < s.screen.height-s.left.height {
		s.left.MoveDown()
	} else {
		s.left.pos.y = s.screen.height - s.left.height
	}
}

func (s *State) RightMoveUp() {
	if s.right.pos.y-s.right.vel > 0 {
		s.right.MoveUp()
	} else {
		s.right.pos.y = 0
	}
}

func (s *State) RightMoveDown() {
	if s.right.pos.y+s.right.vel < s.screen.height-s.right.height {
		s.right.MoveDown()
	} else {
		s.right.pos.y = s.screen.height - s.right.height
	}
}

func (s *State) BallMove() {
	s.ball.move()
}

func (s *State) HandleCollision() {
	// ball hits lower or upper borders
	if s.ball.pos.y <= 0 || s.ball.pos.y >= s.screen.height-s.ball.size {
		s.ball.movec.y *= -1
	} else if (s.ball.pos.x <= s.left.width) && (s.left.pos.y-s.ball.size <= s.ball.pos.y && s.ball.pos.y <= s.left.pos.y+s.left.height) {
		// ball hits left platform
		s.ball.movec.x *= -1
	} else if (s.ball.pos.x >= s.screen.width-s.ball.size-s.right.width) && (s.right.pos.y-s.ball.size <= s.ball.pos.y && s.ball.pos.y <= s.right.pos.y+s.right.height) {
		// ball hits right platform
		s.ball.movec.x *= -1
	} else {
		return
	}

	s.ball.accelerate()
}

func (s *State) HandleOutOfBounds() bool {
	// right scored
	if s.ball.pos.x+s.ball.size < 0 {
		s.score.right++
		return true
	}

	if s.ball.pos.x >= s.right.pos.x+s.right.width {
		// log.Println("XYU")
		s.score.left++
		return true
	}

	return false
}

func (s State) PlayerWon() bool {
	if s.score.left > s.score.max || s.score.right > s.score.max {
		return true
	}

	return false
}

func (s State) MaxScore() int {
	return s.score.max
}

type ball struct {
	size float64
	pos  vector
	// movement vector
	movec vector
	// acceleration multiplier
	accel float64
}

func initBall(size, velX, velY, accel float64) ball {
	rng := rand.IntN(2)

	movec := vector{}
	if rng == 1 {
		movec.x = velX
	} else {
		movec.x = -velX
	}

	rng = rand.IntN(2)
	if rng == 1 {
		movec.y = velY
	} else {
		movec.y = -velY
	}

	return ball{
		size:  size,
		movec: movec,
		accel: accel,
	}
}

// cur position + movement vector
func (b *ball) move() {
	b.pos.add(b.movec)
}

// accelerate ball by mult
func (b *ball) accelerate() {
	b.movec.mult(b.accel)
}

type platform struct {
	width  float64
	height float64
	vel    float64
	pos    vector
}

// move platform either up or down
func (p *platform) MoveUp() {
	p.pos.y -= p.vel
}

func (p *platform) MoveDown() {
	p.pos.y += p.vel
}

type vector struct {
	x float64
	y float64
}

func (v1 *vector) add(v2 vector) {
	v1.x += v2.x
	v1.y += v2.y
}

func (v1 *vector) mult(mult float64) {
	v1.x *= mult
	v1.y *= mult
}

type screen struct {
	width  float64
	height float64
}
