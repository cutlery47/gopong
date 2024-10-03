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
		ball:   initBall(conf.BallSize, conf.BallInitVelX, conf.BallInitVelY),
		left:   platform{width: conf.PlatWidth, height: conf.PlatHeight, vel: conf.PlatVelocity},
		right:  platform{width: conf.PlatWidth, height: conf.PlatHeight, vel: conf.PlatVelocity},
		screen: screen{width: conf.ScreenWidth, height: conf.ScreenHeight},
	}
}

func (s *State) Place() {
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
	if s.ball.pos.y == 0 || s.ball.pos.y == s.screen.height-s.ball.size {
		s.ball.movec.y *= -1
	}

	if 0 <= s.ball.pos.x && s.ball.pos.x <= s.left.width {
		s.ball.movec.x *= -1
	}

	if s.right.pos.x-s.ball.size <= s.ball.pos.x && s.ball.pos.x <= s.screen.width-s.ball.size {
		s.ball.movec.x *= -1
	}

}

type ball struct {
	size float64
	pos  vector
	// movement vector
	movec vector
}

func initBall(size float64, velX float64, velY float64) ball {
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
	}
}

// cur position + movement vector
func (b *ball) move() {
	b.pos.add(b.movec)
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
