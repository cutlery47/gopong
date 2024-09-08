package handlers

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"gopong/server/internal/packet"

	"github.com/gorilla/websocket"
)

type SessionHandler struct {
	sessions []Session
}

type Session struct {
	player1 *websocket.Conn
	player2 *websocket.Conn
}

func (handler *SessionHandler) Handle(sesh Session) {
	log.Println("Started session handling")

	buf1 := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)

	enc1 := gob.NewEncoder(buf1)
	enc2 := gob.NewEncoder(buf2)

	for {
		somedata := packet.Packet{
			State: packet.StateMatchmaking,
			Left: packet.VectorPackage{
				X: 0,
				Y: 0,
			},
			Right: packet.VectorPackage{
				X: 0,
				Y: 0,
			},
			Ball: packet.VectorPackage{
				X: 0,
				Y: 0,
			},
		}

		log.Println("Enc1:", enc1.Encode(somedata))
		log.Println("Enc2:", enc2.Encode(somedata))

		log.Println("buf1:", buf1)
		log.Println("buf2:", buf2)

		log.Println("WebsockErr1:", sesh.player1.WriteMessage(websocket.TextMessage, buf1.Bytes()))
		log.Println("WebsockErr2:", sesh.player2.WriteMessage(websocket.TextMessage, buf2.Bytes()))

		buf1.Reset()
		buf2.Reset()

		time.Sleep(time.Second * 1)
	}
}

func newSessionHandler() *SessionHandler {
	sessions := []Session{}

	return &SessionHandler{
		sessions: sessions,
	}
}
