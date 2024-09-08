package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	qHandler *QueueHandler
	upgrader websocket.Upgrader
}

func (handler *WebsocketHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := handler.upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("ServeHTTP: ", err)
		w.WriteHeader(400)
	}

	handler.qHandler.queue = append(handler.qHandler.queue, conn)
}

func NewWebsocketHandler() *WebsocketHandler {
	sHandler := newSessionHandler()
	qHandler := newQueueHandler(sHandler)
	upgrader := websocket.Upgrader{}

	go qHandler.Handle()

	return &WebsocketHandler{
		qHandler: qHandler,
		upgrader: upgrader,
	}
}
