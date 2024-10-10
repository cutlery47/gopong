package core

type Client struct {
	state *State

	// channel for reading user inputs
	inputChan <-chan CombinedKeyboardInputResult
	// channel for terminating main loop
	exitChan <-chan byte
	// channel for signaling that the game has finished
	finishChan chan<- byte
}

func InitClient(inputChan <-chan CombinedKeyboardInputResult, exitChan <-chan byte, finishChan chan<- byte, state *State) Client {
	state.Flush()

	return Client{
		state:      state,
		inputChan:  inputChan,
		exitChan:   exitChan,
		finishChan: finishChan,
	}
}

func (c *Client) Run() {
	for {
		if exit := c.listenForExit(); exit {
			return
		}

		input := <-c.inputChan

		if input.Left.Up {
			c.state.LeftMoveUp()
		}
		if input.Left.Down {
			c.state.LeftMoveDown()
		}

		if input.Right.Up {
			c.state.RightMoveUp()
		}

		if input.Right.Down {
			c.state.RightMoveDown()
		}

		c.state.BallMove()

		scored := c.state.HandleOutOfBounds()

		if scored {
			if c.state.PlayerWon() {
				c.finishChan <- 1
			}
			c.state.Flush()
		}

		c.state.HandleCollision()
	}
}

func (c Client) listenForExit() bool {
	select {
	case <-c.exitChan:
		return true
	default:
		return false
	}
}

// func (c Client) readInput(timeout time.Duration) (CombinedKeyboardInputResult, error) {
// 	select {
// 	case input := <-c.inputChan:
// 		return input, nil
// 	case <-time.After(timeout):
// 		return CombinedKeyboardInputResult{}, ErrClientReadTimedOut
// 	}
// }
