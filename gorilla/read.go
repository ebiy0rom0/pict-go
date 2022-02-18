package gorilla

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func (c *Client) readPump() {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// 改行 -> SP 置換
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.handler(message)
	}
}
