package game

import (
	"gopong/server/internal/pack"
	"log"

	"github.com/gorilla/websocket"
)

// incoming connection
type connection struct {
	conn *websocket.Conn
}

func (c connection) Send(pack pack.ServerPacket) (err error) {
	err = c.conn.WriteJSON(pack)
	if err != nil {
		log.Println("connection.Send():", err)
		return err
	}
	return err
}

func (c connection) Read(pack *pack.ClientPacket) (err error) {
	err = c.conn.ReadJSON(pack)
	if err != nil {
		log.Println("connection.Read():", err)
		return err
	}
	return err
}

func (c connection) SendStatus(status pack.PlayerStatus) {
	pack := pack.ServerPacket{
		Status: status,
	}
	c.Send(pack)
}

func (c connection) Close() {
	err := c.conn.Close()
	if err != nil {
		log.Println(err)
	}
}
