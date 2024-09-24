package main

import (
	"gopong/client/internal/game"
	"log"
)

func main() {
	client := game.NewMultiplayerClient()
	if client == nil {
		log.Println("Error when initializing a client...")
		return
	}

	game := game.New(client)
	if game == nil {
		log.Println("Error when starting a game...")
		return
	}

	err := game.Run()
	if err != nil {
		log.Println("Runtime error:", err)
		return
	}
}
