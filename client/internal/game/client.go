package game

import (
	"fmt"
	"gopong/client/internal/gui"
	"log"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

type Client interface {
	ebiten.Game
}

type localClient struct {
	updater Updater
	drawer  Drawer
	window  *gui.Window
}

func NewLocalClient() *localClient {
	updater := NewLocalUpdater()
	drawer := NewRenderer()

	ebiten.SetWindowSize(1000, 500)
	window := gui.NewWindow(1000, 500)

	return &localClient{
		updater: updater,
		drawer:  drawer,
		window:  window,
	}
}

func (lc *localClient) Update() error {
	err := lc.updater.Update(lc.window)
	if err != nil {
		return fmt.Errorf("lc.updater.Update: %v", err)
	}
	return nil
}

func (lc *localClient) Draw(screen *ebiten.Image) {
	lc.drawer.Draw(lc.window, screen)
}

func (lc *localClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

type connection struct {
	conn *websocket.Conn
}

type multiplayerClient struct {
	conn connection
}

func NewMultiplayerClient() *multiplayerClient {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Println("Couldn't dial:", err)
		return nil
	}

	client := &multiplayerClient{
		conn: connection{conn: conn},
	}

	return client
}

func (mc *multiplayerClient) Update() error {
	log.Println("update")
	return nil
}

func (mc *multiplayerClient) Draw(screen *ebiten.Image) {
	log.Println("draw")
}

func (mc *multiplayerClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
