package main

import (
	"gopong/client/internal/game"
	"log"
)

func main() {
	if err := game.Run("./config/config.yaml"); err != nil {
		log.Fatal(err)
	}
}
