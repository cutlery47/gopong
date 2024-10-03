package local

import (
	"log"
	"time"

	"github.com/cutlery47/gopong/client/config"
	"github.com/cutlery47/gopong/client/internal/core"
)

// client for local play
type LocalClient struct {
	updateChan chan<- core.StateUpdate
	inputChan  <-chan core.CombinedKeyboardInputResult
}

func InitLocalClient(updateChan chan<- core.StateUpdate, inputChan <-chan core.CombinedKeyboardInputResult, conf config.LocalConfig) LocalClient {
	return LocalClient{
		updateChan: updateChan,
		inputChan:  inputChan,
	}
}

func Init() LocalClient {
	return LocalClient{}
}

func (lc LocalClient) Run() {
	exitChan := make(chan byte)
	go lc.Update(exitChan)
	go lc.Read(exitChan)
	for {
		time.Sleep(time.Second)
	}

}

func (lc LocalClient) Update(exitChan <-chan byte) {
	for {
		lc.updateChan <- core.StateUpdate{LeftOffset: 1.0, RightOffset: 2.0, BallOffsetX: 0.5, BallOffsetY: 0.5}
	}
}

func (lc LocalClient) Read(exitChan <-chan byte) {
	for {
		input := <-lc.inputChan
		log.Println(input)
	}
}

// func NewClient(cliConfig config.ClientConfig) *localClient {
// 	updater := NewLocalUpdater()
// 	drawer := common.NewRenderer()

// 	ebiten.SetTPS(cliConfig.MaxTPS)

// 	ebiten.SetWindowSize(1000, 500)
// 	canvas := gui.NewCanvas(1000, 500)

// 	return &localClient{
// 		updater: updater,
// 		drawer:  drawer,
// 		canvas:  canvas,
// 	}
// }

// func (lc *localClient) Update() error {
// 	err := lc.updater.Update(lc.canvas)
// 	if err != nil {
// 		return fmt.Errorf("lc.updater.Update: %v", err)
// 	}
// 	return nil
// }

// func (lc *localClient) Draw(screen *ebiten.Image) {
// 	lc.drawer.Draw(lc.canvas, screen)
// }

// func (lc *localClient) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return outsideWidth, outsideHeight
// }
