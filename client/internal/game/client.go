package game

import (
	"fmt"

	"github.com/cutlery47/gopong/client/internal/gui"
	"github.com/cutlery47/gopong/common/conn"
	"github.com/cutlery47/gopong/common/protocol"

	"github.com/hajimehoshi/ebiten/v2"
)

type Client interface {
	ebiten.Game
}

type localClient struct {
	updater *localUpdater
	drawer  Drawer
	canvas  *gui.Canvas
}

func NewLocalClient() *localClient {
	updater := NewLocalUpdater()
	drawer := NewRenderer()

	ebiten.SetWindowSize(1000, 500)
	canvas := gui.NewCanvas(1000, 500)

	return &localClient{
		updater: updater,
		drawer:  drawer,
		canvas:  canvas,
	}
}

func (lc *localClient) Update() error {
	err := lc.updater.Update(lc.canvas)
	if err != nil {
		return fmt.Errorf("lc.updater.Update: %v", err)
	}
	return nil
}

func (lc *localClient) Draw(screen *ebiten.Image) {
	lc.drawer.Draw(lc.canvas, screen)
}

func (lc *localClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

type multiplayerClient struct {
	conn      conn.Connection
	canvas    *gui.Canvas
	drawer    Drawer
	reader    KeyboardInputReader
	side      string
	statePipe <-chan protocol.ServerPacket
	inputPipe <-chan KeyboardInputResult
}

func NewMultiplayerClient(conn conn.Connection, config protocol.GameConfig) *multiplayerClient {
	inputPipe := make(chan KeyboardInputResult)
	statePipe := make(chan protocol.ServerPacket)

	canvas := gui.NewCanvasFromConfig(config)
	ebiten.SetWindowSize(int(config.CanvasWidth), int(config.CanvasHeight))

	var reader KeyboardInputReader
	if config.Side == "left" {
		reader = KeyboardInputReader{
			upKey:   ebiten.KeyW,
			downKey: ebiten.KeyS,
		}
	} else {
		reader = KeyboardInputReader{
			upKey:   ebiten.KeyArrowUp,
			downKey: ebiten.KeyArrowDown,
		}
	}

	client := &multiplayerClient{
		conn:      conn,
		canvas:    canvas,
		drawer:    NewRenderer(),
		reader:    reader,
		side:      string(config.Side),
		statePipe: statePipe,
		inputPipe: inputPipe,
	}

	go conn.ListenFromServer(statePipe)

	return client
}

// this is where game state updates
// !!!this should probably only consume server data and update gui elements accordingly!!!
func (mc *multiplayerClient) Update() error {
	newState := <-mc.statePipe
	mc.canvas.Update(newState)

	pack := protocol.ClientPacket{}
	if mc.side == "left" {
		pack.Position = protocol.Vector{X: mc.canvas.Left.XCoord(), Y: mc.canvas.Left.YCoord()}
	} else {
		pack.Position = protocol.Vector{X: mc.canvas.Right.XCoord(), Y: mc.canvas.Right.YCoord()}
	}

	input := mc.reader.Read()
	if input.up {
		pack.Position.Y -= 5
	}

	if input.down {
		pack.Position.Y += 5
	}

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
