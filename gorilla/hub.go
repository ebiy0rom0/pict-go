package gorilla

import (
	"context"
	"sync"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan map[*Client][]byte
}

var hub *Hub
var once sync.Once

func RunHub(ctx context.Context) {
	once.Do(func() {
		hub = newHub()
		go hub.run(ctx)
	})
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run(c context.Context) {
	for {
		select {
		case <-c.Done():
			close(h.register)
			close(h.unregister)
			return

		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			delete(h.clients, client)
			close(client.send)

		case broadcast := <-h.broadcast:
			for client := range h.clients {
				message, ok := broadcast[client]
				if ok {
					continue
				}

				client.send <- message

				close(client.send)
				h.unregister <- client
			}
		}
	}
}
