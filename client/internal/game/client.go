package game

import (
	"gopong/client/config"
	"gopong/client/internal/gui"
)

type ClientUpdater struct {
	game GameUpdater
}

func (u ClientUpdater) Update(cfg *config.Config, left *gui.Platform, right *gui.Platform, ball *gui.Ball) error {
	return nil
}
