package local

import (
	"fmt"

	"github.com/cutlery47/gopong/client/internal/game/common"
	"github.com/cutlery47/gopong/client/internal/gui"
	"github.com/hajimehoshi/ebiten/v2"
)

type localClient struct {
	drawer  common.Drawer
	updater common.Updater
	canvas  *gui.Canvas
}

func NewClient() *localClient {
	updater := NewLocalUpdater()
	drawer := common.NewRenderer()

	ebiten.SetWindowSize(1000, 500)
	canvas := gui.NewCanvas(1000, 500)

	return &localClient{
		updater: updater,
		drawer:  drawer,
		canvas:  canvas,
	}
}

func (lc *localClient) Update() error {
	err := lc.updater.Update(lc.canvas)
	if err != nil {
		return fmt.Errorf("lc.updater.Update: %v", err)
	}
	return nil
}

func (lc *localClient) Draw(screen *ebiten.Image) {
	lc.drawer.Draw(lc.canvas, screen)
}

func (lc *localClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
