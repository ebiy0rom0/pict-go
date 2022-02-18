package gorilla

// メッセージハンドラ
func (c *Client) handler(message []byte) {
	hub.broadcast <- map[*Client][]byte{
		c: message,
	}
}
