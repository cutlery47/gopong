package core

type State struct {
	ball  *ball
	left  *platform
	right *platform
}

func NewState(ball *ball, left, right *platform) *State {
	return &State{
		ball:  ball,
		left:  left,
		right: right,
	}
}

// entity movement
type StateUpdate struct {
	LeftOffset  float64
	RightOffset float64
	BallOffsetX float64
	BallOffsetY float64
}

func (st *State) Update(upd StateUpdate) {
	// firstly we update platform positions
	st.left.move(upd.LeftOffset)
	st.right.move(upd.RightOffset)

	// then we update ball position
	st.ball.move(vector{x: upd.BallOffsetX, y: upd.BallOffsetY})
}

type ball struct {
	pos vector
}

// cur position + movement vector
func (b *ball) move(movec vector) {
	b.pos.add(movec)
}

type platform struct {
	pos vector
}

// move platform either up or down
func (p *platform) move(offset float64) {
	p.pos.y += offset
}

type vector struct {
	x float64
	y float64
}

func (v1 *vector) add(v2 vector) {
	v1.x += v2.x
	v1.y += v2.y
}

// func (v1 *vector) mult(mult float64) {
// 	v1.x *= mult
// 	v1.y *= mult
// }
