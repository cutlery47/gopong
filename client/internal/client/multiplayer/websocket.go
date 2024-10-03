package multiplayer

import (
	"fmt"
	"time"

	"github.com/cutlery47/gopong/common/protocol"
	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	conn     *websocket.Conn
	exitChan chan<- byte
}

func NewWebsocketClient(host, port string, exitChan chan<- byte) (*WebsocketClient, error) {
	url := fmt.Sprintf("ws://%v:%v", host, port)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &WebsocketClient{conn: conn, exitChan: exitChan}, nil
}

func (wc WebsocketClient) Read(buff *protocol.ServerPacket) error {
	err := wc.conn.ReadJSON(buff)
	if err != nil {
		if websocket.IsCloseError(err) {
			wc.Close()
		}
		return err
	}
	return nil
}

func (wc WebsocketClient) Write(data protocol.ClientPacket) error {
	err := wc.conn.WriteJSON(data)
	if err != nil {
		if websocket.IsCloseError(err) {
			wc.Close()
		}
		return err
	}
	return nil
}

func (wc WebsocketClient) Close() error {
	writeDeadline := time.Now().Add(time.Minute)
	readDeadline := time.Now().Add(5 * time.Second)

	// signaling that the connection has been sold
	wc.exitChan <- 1

	// cleaning up
	if err := wc.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		writeDeadline,
	); err != nil && err != websocket.ErrCloseSent {
		return wc.Close()
	}

	wc.conn.SetReadDeadline(readDeadline)
	for {
		_, _, err := wc.conn.NextReader()
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			break
		}
		if err != nil {
			break
		}
	}

	return wc.conn.Close()
}
