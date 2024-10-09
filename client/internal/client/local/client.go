package local

import (
	"log"

	"github.com/cutlery47/gopong/client/internal/core"
)

// client for local play
type LocalClient struct {
	state     *core.State
	inputChan <-chan core.CombinedKeyboardInputResult
	exitChan  chan<- byte
}

func InitClient(inputChan <-chan core.CombinedKeyboardInputResult, exitChan chan<- byte, state *core.State) LocalClient {
	state.Flush()

	return LocalClient{
		inputChan: inputChan,
		exitChan:  exitChan,
		state:     state,
	}
}

func (lc *LocalClient) Run() {
	initialState := *lc.state
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

		scored := lc.state.HandleOutOfBounds()

		if scored {
			log.Println("here")
			if lc.state.PlayerWon() {
				lc.exitChan <- 1
				return
			}
			*lc.state = initialState

		}

		lc.state.HandleCollision()
	}
}
