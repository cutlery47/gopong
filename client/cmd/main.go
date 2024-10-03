package main

import (
	game "github.com/cutlery47/gopong/client/internal"
)

func main() {
	game := game.NewLocalGame("config/config.json")
	game.Run()
}
