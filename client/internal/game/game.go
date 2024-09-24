package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	client Client
}

func New(client Client) *Game {
	return &Game{
		client: client,
	}
}

func (g *Game) Run() error {
	err := ebiten.RunGame(g.client)
	if err != nil {
		return fmt.Errorf("Game.Run(): %v", err)
	}
	return nil
}
