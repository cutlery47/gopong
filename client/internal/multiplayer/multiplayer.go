package multiplayer

import (
	"github.com/cutlery47/gopong/client/config"
	"github.com/cutlery47/gopong/client/internal/core"
)

type Multiplayer struct {
	// channel for reading session input
	sessionInputChan <-chan core.CombinedKeyboardGameInputResult
	// channel for sending output to client
	clientOutputChan chan<- core.CombinedKeyboardGameInputResult
	// channel for terminating main loop
	exitChan <-chan byte
	// channel for signaling that multiplayer client is about to close
	finishChan chan<- byte
}

func Init(
	config config.Config,
	sessionInputChan <-chan core.CombinedKeyboardGameInputResult,
	clientOutputChan chan<- core.CombinedKeyboardGameInputResult,
	exitChan <-chan byte,
	finishChan <-chan byte,
) Multiplayer {
	return Multiplayer{}
}

func (m Multiplayer) FindGame() *core.State {
	return &core.State{}
}

func (m Multiplayer) Run() {
	for {
		input := <-m.sessionInputChan

		output := input

		m.clientOutputChan <- output
	}
}

func (m Multiplayer) listenForExit() bool {
	select {
	case <-m.exitChan:
		return true
	default:
		return false
	}
}

// import (
// 	"github.com/cutlery47/gopong/client/config"
// 	"github.com/cutlery47/gopong/client/internal/game/common"
// 	"github.com/cutlery47/gopong/client/internal/gui"
// 	"github.com/cutlery47/gopong/common/conn"
// 	"github.com/cutlery47/gopong/common/protocol"

// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type multiplayerClient struct {
// 	updater *Updater
// 	drawer  common.Drawer

// 	side   string
// 	conn   conn.Connection
// 	canvas *gui.Canvas

// 	inputPipe <-chan common.KeyboardInputResult
// }

// func NewMultiplayerClient(conn conn.Connection, servConfig protocol.GameConfig, cliConfig config.GameConfig) *multiplayerClient {
// 	inputPipe := make(chan common.KeyboardInputResult)
// 	statePipe := make(chan protocol.ServerPacket)

// 	var reader common.KeyboardInputReader
// 	if servConfig.Side == "left" {
// 		reader = common.LeftKeyboardInputReader
// 	} else {
// 		reader = common.RightKeyboardInputReader
// 	}

// 	ebiten.SetTPS(cliConfig.MaxTPS)

// 	updater := NewUpdater(statePipe, &reader, string(servConfig.Side))
// 	drawer := common.NewRenderer()
// 	side := string(servConfig.Side)

// 	canvas := gui.NewCanvasFromConfig(servConfig)
// 	ebiten.SetWindowSize(int(servConfig.CanvasWidth), int(servConfig.CanvasHeight))

// 	client := &multiplayerClient{
// 		updater:   updater,
// 		drawer:    drawer,
// 		conn:      conn,
// 		canvas:    canvas,
// 		side:      side,
// 		inputPipe: inputPipe,
// 	}

// 	go conn.ListenFromServer(statePipe)

// 	return client
// }

// func (mc *multiplayerClient) Update() error {
// 	mc.updater.Update(mc.canvas)
// 	pack := mc.updater.PackState(mc.canvas)
// 	mc.conn.Send(pack)
// 	return nil
// }

// // this is where game state in rendered
// func (mc *multiplayerClient) Draw(screen *ebiten.Image) {
// 	mc.drawer.Draw(mc.canvas, screen)
// }

// // this is bs
// func (mc *multiplayerClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return outsideWidth, outsideHeight
// }
