package common

import (
	"github.com/cutlery47/gopong/client/internal/gui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Client interface {
	ebiten.Game
}

type Updater interface {
	Update(canvas *gui.Canvas) error
}

type Drawer interface {
	Draw(window *gui.Canvas, screen *ebiten.Image)
}
