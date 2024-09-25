package conn

import (
	"log"

	"github.com/cutlery47/gopong/common/protocol"
	"github.com/gorilla/websocket"
)

// incoming connection
type Connection struct {
	// underlying websocket
	conn *websocket.Conn
	// channel for notifying that connection has been closed
	removeConnPipe <-chan byte
}

func New(wc *websocket.Conn) Connection {
	pipe := make(chan byte)

	// when Close() call is received -> send some data to the pipe
	// so that anyone who keeps track of the connection knows, that it has been closed
	wc.SetCloseHandler(
		func(code int, text string) error {
			pipe <- 1
			wc.Close()
			return nil
		},
	)

	return Connection{
		conn:           wc,
		removeConnPipe: pipe,
	}
}

func (c Connection) Send(pack protocol.ServerPacket) (err error) {
	err = c.conn.WriteJSON(pack)
	if err != nil {
		log.Println("connection.Send():", err)
		return err
	}
	return err
}

func (c Connection) Read(pack *protocol.ClientPacket) (err error) {
	err = c.conn.ReadJSON(pack)
	if err != nil {
		log.Println("connection.Read():", err)
		return err
	}
	return err
}

func (c Connection) SendStatus(status protocol.PlayerStatus) {
	pack := protocol.ServerPacket{
		Status: status,
	}
	c.Send(pack)
}

func (c Connection) Close() {
	err := c.conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func (c Connection) Pipe() <-chan byte {
	return c.removeConnPipe
}

func (c Connection) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}
