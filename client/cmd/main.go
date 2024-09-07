package main

import (
	"gopong/internal/game"
	"log"
)

func main() {
	if err := game.Run("./config/config.yaml"); err != nil {
		log.Fatal(err)
	}
}
