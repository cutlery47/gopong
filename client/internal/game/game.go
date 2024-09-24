package game

import (
	"log"

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

func (g *Game) Run() {
	err := ebiten.RunGame(g.client)
	if err != nil {
		log.Printf("Game.Run(): %v", err)
	}
}

// type Game struct {
// 	updater Updater
// 	drawer  Drawer
// 	window  *gui.Window
// 	config  *config.Config

// 	client Client
// }

// func newGame(configPath string) (*Game, error) {
// 	game := &Game{}

// 	client := NewClient()
// 	if client == nil {
// 		return nil, fmt.Errorf("NewClient: Couldn't create a client")
// 	}

// 	config, err := config.New(configPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("newGame: %w", err)
// 	}

// 	window := gui.NewWindow(config.Window.Width, config.Window.Height)
// 	ebiten.SetWindowSize(config.Window.Width, config.Window.Height)

// 	game.config = config
// 	game.window = window
// 	game.updater = *newLocalUpdater()
// 	game.drawer = GameDrawer{}

// 	return game, nil
// }

// func Run(configPath string) error {
// 	game, err := newGame(configPath)
// 	if err != nil {
// 		return fmt.Errorf("Run: %w", err)
// 	}

// 	err = ebiten.RunGame(game)
// 	if err != nil {
// 		return fmt.Errorf("Run: %w", err)
// 	}

// 	return nil
// // }

// // func (g *Game) Update() error {
// // 	err := g.updater.Update(g.config, g.window.Platform_1, g.window.Platform_2, g.window.Ball)
// // 	if err != nil {
// // 		return fmt.Errorf("Update: %w", err)
// // 	}

// // 	return nil
// // }

// // func (g *Game) Draw(screen *ebiten.Image) {
// // 	g.drawer.Draw(g.window.Platform_1, g.window.Platform_2, g.window.Ball, screen)
// // }

// // func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// // 	return outsideWidth, outsideHeight
// // }
