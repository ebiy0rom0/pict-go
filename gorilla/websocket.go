package gorilla

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func ServeWS(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := createClient(conn)
	hub.register <- c

	go c.readPump()
	go c.writePump()
}
