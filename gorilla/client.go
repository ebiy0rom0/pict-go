package gorilla

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	room int
}

// クライアント 作成
func createClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}
}
