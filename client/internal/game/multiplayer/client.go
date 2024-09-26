package multiplayer

import (
	"github.com/cutlery47/gopong/client/internal/game/common"
	"github.com/cutlery47/gopong/client/internal/gui"
	"github.com/cutlery47/gopong/common/conn"
	"github.com/cutlery47/gopong/common/protocol"

	"github.com/hajimehoshi/ebiten/v2"
)

type multiplayerClient struct {
	updater *Updater
	drawer  common.Drawer
	reader  common.KeyboardInputReader

	side   string
	conn   conn.Connection
	canvas *gui.Canvas

	inputPipe <-chan common.KeyboardInputResult
}

func NewMultiplayerClient(conn conn.Connection, config protocol.GameConfig) *multiplayerClient {
	inputPipe := make(chan common.KeyboardInputResult)
	statePipe := make(chan protocol.ServerPacket)

	var reader common.KeyboardInputReader
	if config.Side == "left" {
		reader = common.LeftKeyboardInputReader
	} else {
		reader = common.RightKeyboardInputReader
	}

	updater := NewUpdater(statePipe, &reader)
	drawer := common.NewRenderer()
	side := string(config.Side)

	canvas := gui.NewCanvasFromConfig(config)
	ebiten.SetWindowSize(int(config.CanvasWidth), int(config.CanvasHeight))

	client := &multiplayerClient{
		updater:   updater,
		reader:    reader,
		drawer:    drawer,
		conn:      conn,
		canvas:    canvas,
		side:      side,
		inputPipe: inputPipe,
	}

	go conn.ListenFromServer(statePipe)

	return client
}

func (mc *multiplayerClient) Update() error {
	mc.updater.Update(mc.canvas)
	pack := mc.updater.PackState(mc.side, mc.canvas)
	mc.conn.Send(pack)
	return nil
}

// this is where game state in rendered
func (mc *multiplayerClient) Draw(screen *ebiten.Image) {
	mc.drawer.Draw(mc.canvas, screen)
}

// this is bs
func (mc *multiplayerClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
