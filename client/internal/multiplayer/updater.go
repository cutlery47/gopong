package multiplayer

// import (
// 	"github.com/cutlery47/gopong/client/internal/game/common"
// 	"github.com/cutlery47/gopong/client/internal/gui"
// 	"github.com/cutlery47/gopong/common/protocol"
// )

// type Updater struct {
// 	reader *common.KeyboardInputReader

// 	statePipe <-chan protocol.ServerPacket
// }

// func NewUpdater(statePipe <-chan protocol.ServerPacket, reader *common.KeyboardInputReader, side string) *Updater {
// 	return &Updater{
// 		statePipe: statePipe,
// 		reader:    reader,
// 	}
// }

// func (u *Updater) Update(canvas *gui.Canvas) error {
// 	newState := <-u.statePipe
// 	u.updateCanvasWithState(canvas, newState)
// 	return nil
// }

// func (u *Updater) PackState(canvas *gui.Canvas) protocol.ClientPacket {
// 	pack := protocol.ClientPacket{}

// 	input := u.reader.Read()
// 	if input.Up {
// 		pack.InputUp = true
// 	}
// 	if input.Down {
// 		pack.InputDown = true
// 	}

// 	return pack
// }

// func (u *Updater) updateCanvasWithState(canvas *gui.Canvas, state protocol.ServerPacket) {
// 	canvas.Left().SetPosition(state.State.LeftPosition.X, state.State.LeftPosition.Y)
// 	canvas.Right().SetPosition(state.State.RightPosition.X, state.State.RightPosition.Y)
// 	canvas.Ball().SetPosition(state.State.BallPosition.X, state.State.BallPosition.Y)
// }
