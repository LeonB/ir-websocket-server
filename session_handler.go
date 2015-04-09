package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr     = ":8080"
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func newSessionHandler() func(http.ResponseWriter, *http.Request) {
	handler := sessionHandler{}
	return handler.serve
}

type sessionHandler struct {
	ws *websocket.Conn
}

func (h *sessionHandler) serve(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	h.ws = ws

	go h.writer()
	h.reader()
}

func (h *sessionHandler) writer() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			p := []byte("jemoeder")
			if err := h.ws.WriteMessage(websocket.TextMessage, p); err != nil {
				return
			}
		}
	}
}

func (h *sessionHandler) reader() {
	defer h.ws.Close()

	// Time allowed to read the next pong message from the client.
	pongWait := 60 * time.Second

	h.ws.SetReadLimit(512)
	h.ws.SetReadDeadline(time.Now().Add(pongWait))
	h.ws.SetPongHandler(func(string) error {
		h.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// Implement pull commands
	for {
		_, p, err := h.ws.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println(string(p))
	}
}
