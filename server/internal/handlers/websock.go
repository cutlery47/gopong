package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	queueHandler *QueueHandler
	upgrader     websocket.Upgrader
}

func (handler *WebsocketHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := handler.upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("ServeHTTP: ", err)
		w.WriteHeader(400)
	}

	handler.queueHandler.queue = append(handler.queueHandler.queue, conn)
}

func NewWebsocketHandler() *WebsocketHandler {
	gameHandler := new(GameHandler)
	configHandler := new(ConfigHandler)

	sessionHandler := newSessionHandler(
		gameHandler,
	)

	queueHandler := newQueueHandler(
		sessionHandler,
		configHandler,
	)

	upgrader := websocket.Upgrader{}

	go queueHandler.Handle()

	return &WebsocketHandler{
		queueHandler: queueHandler,
		upgrader:     upgrader,
	}
}
