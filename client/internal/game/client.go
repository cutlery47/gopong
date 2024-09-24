package game

import (
	"fmt"
	"gopong/client/internal/gui"
	"gopong/client/internal/pack"
	"log"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

type Client interface {
	ebiten.Game
}

type localClient struct {
	updater *localUpdater
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
	conn      *websocket.Conn
	statePipe chan<- pack.ServerPacket
}

func InitConnection(url string, pipe chan<- pack.ServerPacket) (connection, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("InitConnection: %v", err)
		return connection{}, err
	}

	return connection{conn: conn, statePipe: pipe}, nil
}

func (c connection) Read(buff *pack.ServerPacket) error {
	err := c.conn.ReadJSON(buff)
	if err != nil {
		return err
	}
	return nil
}

func (c connection) Send(pack pack.ClientPacket) (err error) {
	err = c.conn.WriteJSON(pack)
	if err != nil {
		log.Println("connection.Send():", err)
		return err
	}
	return err
}

func (c connection) SendACK() (err error) {
	err = c.conn.WriteJSON(pack.ClientAck)
	if err != nil {
		log.Println("connection.SendACK():", err)
		return err
	}
	return err
}

func (c connection) Listen() {
	for {
		state := pack.ServerPacket{}
		c.Read(&state)
		log.Println("SENDING")
		c.statePipe <- state
	}
}

type multiplayerClient struct {
	conn      connection
	window    *gui.Window
	drawer    Drawer
	statePipe <-chan pack.ServerPacket
	inputPipe <-chan KeyboardInputResult
}

func NewMultiplayerClient() *multiplayerClient {
	inputPipe := make(chan KeyboardInputResult)
	statePipe := make(chan pack.ServerPacket)

	conn, err := InitConnection("ws://localhost:8080", statePipe)
	if err != nil {
		return nil
	}

	window := gui.NewWindow(1000, 500)
	ebiten.SetWindowSize(1000, 500)

	client := &multiplayerClient{
		conn:      conn,
		window:    window,
		drawer:    NewRenderer(),
		statePipe: statePipe,
		inputPipe: inputPipe,
	}

	go conn.Listen()
	go client.HandleInput()

	return client
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
	mc.window.Update(newState)

	return nil
}

// this is where game state in rendered
func (mc *multiplayerClient) Draw(screen *ebiten.Image) {
	mc.drawer.Draw(mc.window, screen)
}

// this is bs
func (mc *multiplayerClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
