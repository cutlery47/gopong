package game

import (
	"gopong/server/internal/entities"
	"gopong/server/internal/pack"
)

type State struct {
	left  *entities.Platform
	right *entities.Platform
	ball  *entities.Ball
}

func InitState() *State {
	ball := entities.NewBall(
		10,
		entities.Vector{X: 100, Y: 100},
	)

	left := entities.NewPlatform(
		100,
		200,
		entities.Vector{X: 500, Y: 500},
	)

	right := entities.NewPlatform(
		400,
		500,
		entities.Vector{X: 300, Y: 300},
	)

	return &State{
		left:  left,
		right: right,
		ball:  ball,
	}
}

func (s *State) Update(leftInput, rightInput pack.ClientPacket) {

}
