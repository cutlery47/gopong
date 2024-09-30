package conn

import (
	"errors"
	"fmt"
	"log"
	"time"

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

	conn := Connection{
		conn:           wc,
		removeConnPipe: pipe,
	}

	return conn
}

func InitConnection(host string, port string) (Connection, protocol.GameConfig, error) {
	url := fmt.Sprintf("ws://%v:%v", host, port)

	// connecting
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("Dial: %v", err)
		return Connection{}, protocol.GameConfig{}, err
	}

	conn := Connection{conn: c}

	// waiting for game
	res := protocol.ServerPacket{}
	for res.Status != protocol.FoundStatus {
		log.Println("In queue...")
		conn.Read(&res)

	}
	conn.Send("ack")

	log.Println("Game found!")

	// receiving config
	config := protocol.GameConfig{}
	conn.Read(&config)
	conn.Send("ack")

	return conn, config, err
}

func (c Connection) Send(data interface{}) (err error) {
	err = c.conn.WriteJSON(data)
	if err != nil {
		log.Println("connection.Send():", err)
		return err
	}
	return err
}

func (c Connection) Read(buff interface{}) (err error) {
	err = c.conn.ReadJSON(buff)
	if err != nil {
		log.Println("Connection.Read():", err)
		return err
	}
	return err
}

func (c Connection) ReadACK() (err error) {
	ackRes := ""
	c.conn.ReadJSON(&ackRes)
	if ackRes != "ack" {
		return errors.New("ack was not received")
	}
	return nil
}

func (c Connection) Close() {
	err := c.conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func (c Connection) RemoveConnPipe() <-chan byte {
	return c.removeConnPipe
}

func (c Connection) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c Connection) ListenFromServer(statePipe chan<- protocol.ServerPacket) {
	cnt := 0
	ticker := time.NewTicker(1 * time.Second)
	for {
		data := protocol.ServerPacket{}
		c.Read(&data)
		statePipe <- data

		select {
		case <-ticker.C:
			log.Println("Received packets:", cnt)
			cnt = 0
		default:
			cnt += 1
		}
	}
}

type Pair struct {
	Conn1 Connection
	Conn2 Connection
}
