package game

import "log"

type Drawer interface {
	Draw() error
}

type Renderer struct{}

func (r Renderer) Draw() error {
	log.Println("draw")
	return nil
}
