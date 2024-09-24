package game

import (
	"fmt"
	"gopong/client/internal/gui"
	"log"

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
	return &localClient{
		updater: NewDefaultUpdater(),
		window:  gui.NewWindow(100, 100),
	}
}

func (lc *localClient) Update() error {
	err := lc.updater.Update()
	if err != nil {
		return fmt.Errorf("lc.updater.Update: %v", err)
	}
	return nil
}

func (lc *localClient) Draw(screen *ebiten.Image) {
	err := lc.drawer.Draw()
	if err != nil {
		log.Println("lc.drawer.Draw():", err)
	}
}

func (lc *localClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// func (g *Game) Draw(screen *ebiten.Image) {
// 	g.drawer.Draw(g.window.Platform_1, g.window.Platform_2, g.window.Ball, screen)
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return outsideWidth, outsideHeight
// }

// type MultiplayerClient struct {
// 	conn *websocket.Conn
// }

// func NewClient() *Client {
// 	dialer := websocket.DefaultDialer
// 	conn, _, err := dialer.Dial("ws://localhost:8080", nil)
// 	if err != nil {
// 		log.Println(err)
// 		return nil
// 	}

// 	client := &Client{
// 		conn: conn,
// 	}

// 	return client
// }
