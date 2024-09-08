package app

import (
	"gopong/server/internal/handlers"
	"gopong/server/internal/router"
	"gopong/server/internal/server"
	"time"
)

func Run(configPath string) {
	handler := handlers.NewWebsocketHandler()

	router := router.NewRouter(handler)

	s := server.New(router, server.ReadTimeout(time.Second*5))
	s.Run()
}
