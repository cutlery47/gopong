package game

import (
	"fmt"
	"log"

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
	statePipe <-chan protocol.ServerPacket
	inputPipe <-chan KeyboardInputResult
}

func NewMultiplayerClient(conn conn.Connection) *multiplayerClient {
	// inputPipe := make(chan KeyboardInputResult)
	// statePipe := make(chan protocol.ServerPacket)

	// conn, err := InitConnection("ws://localhost:8080", statePipe)
	// if err != nil {
	// 	return nil
	// }

	// canvas := gui.NewWindow(1000, 500)
	// ebiten.SetWindowSize(1000, 500)

	// client := &multiplayerClient{
	// 	conn:      conn,
	// 	canvas:    canvas,
	// 	drawer:    NewRenderer(),
	// 	statePipe: statePipe,
	// 	inputPipe: inputPipe,
	// }

	// go conn.Listen()
	// go client.HandleInput()

	// return client
	return nil
}

func (mc *multiplayerClient) HandleInput() {
	// for {
	// 	// input := <-mc.inputPipe

	// }
}

// this is where game state updates
// !!!this should probably only consume server data and update gui elements accordingly!!!
func (mc *multiplayerClient) Update() error {
	newState := <-mc.statePipe
	log.Println(newState)
	mc.canvas.Update(newState)

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
