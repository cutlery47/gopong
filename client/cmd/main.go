package main

import (
	"gopong/client/internal/game"
	"log"
)

func main() {
	client := game.NewLocalClient()

	game := game.New(client)
	if game == nil {
		log.Println("Error when starting a game...")
	}

	err := game.Run()
	if err != nil {
		log.Println("Runtime error:", err)
	}
}
