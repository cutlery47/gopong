package local

import (
	"github.com/cutlery47/gopong/client/internal/core"
)

// client for local play
type LocalClient struct {
	state     *core.State
	inputChan <-chan core.CombinedKeyboardInputResult
}

func InitClient(inputChan <-chan core.CombinedKeyboardInputResult, state *core.State) LocalClient {
	state.Place()

	return LocalClient{
		inputChan: inputChan,
		state:     state,
	}
}

func (lc *LocalClient) Run() {
	for {
		input := <-lc.inputChan

		if input.Left.Up {
			lc.state.LeftMoveUp()
		}
		if input.Left.Down {
			lc.state.LeftMoveDown()
		}

		if input.Right.Up {
			lc.state.RightMoveUp()
		}

		if input.Right.Down {
			lc.state.RightMoveDown()
		}

		lc.state.BallMove()

		lc.state.HandleCollision()
	}
}
