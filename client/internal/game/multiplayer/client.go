package multiplayer

import (
	"github.com/cutlery47/gopong/client/config"
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

func NewMultiplayerClient(conn conn.Connection, servConfig protocol.GameConfig, cliConfig config.GameConfig) *multiplayerClient {
	inputPipe := make(chan common.KeyboardInputResult)
	statePipe := make(chan protocol.ServerPacket)

	var reader common.KeyboardInputReader
	if servConfig.Side == "left" {
		reader = common.LeftKeyboardInputReader
	} else {
		reader = common.RightKeyboardInputReader
	}

	ebiten.SetTPS(cliConfig.MaxTPS)

	updater := NewUpdater(statePipe, &reader)
	drawer := common.NewRenderer()
	side := string(servConfig.Side)

	canvas := gui.NewCanvasFromConfig(servConfig)
	ebiten.SetWindowSize(int(servConfig.CanvasWidth), int(servConfig.CanvasHeight))

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
