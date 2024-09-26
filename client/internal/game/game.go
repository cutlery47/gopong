package game

import (
	"log"

	"github.com/cutlery47/gopong/common/conn"
	"github.com/hajimehoshi/ebiten/v2"
)

func RunLocalGame() {
	client := NewLocalClient()
	err := ebiten.RunGame(client)
	if err != nil {
		log.Printf("A runtime error occurred: %v", err)
	}
}

func RunMultiplayerGame() {
	conn, config, err := conn.InitConnection("ws://localhost:8080")
	if err != nil {
		log.Println("Couldn't establish connection with the server...")
		return
	}

	client := NewMultiplayerClient(conn, config)
	err = ebiten.RunGame(client)
	if err != nil {
		log.Printf("A runtime error occurred: %v", err)
		return
	}
}
