package game

import (
	"fmt"

	"gopong/client/config"
	"gopong/client/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	window  *gui.Window
	updater Updater
	drawer  Drawer
	config  *config.Config
}

func (g *Game) Update() error {
	err := g.updater.Update(g.config, g.window.Platform_1, g.window.Platform_2, g.window.Ball)
	if err != nil {
		return fmt.Errorf("Update: %w", err)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawer.Draw(g.window.Platform_1, g.window.Platform_2, g.window.Ball, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func newGame(configPath string) (ebiten.Game, error) {
	game := &Game{}

	config, err := config.New(configPath)
	if err != nil {
		return nil, fmt.Errorf("newGame: %w", err)
	}

	window := gui.NewWindow(config.Window.Width, config.Window.Height)
	ebiten.SetWindowSize(config.Window.Width, config.Window.Height)

	game.config = config
	game.window = window
	game.updater = *newLocalUpdater()
	game.drawer = GameDrawer{}

	return game, nil
}

func Run(configPath string) error {
	game, err := newGame(configPath)
	if err != nil {
		return fmt.Errorf("Run: %w", err)
	}

	err = ebiten.RunGame(game)
	if err != nil {
		return fmt.Errorf("Run: %w", err)
	}

	return nil
}
