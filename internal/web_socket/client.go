package web_socket

import "github.com/gorilla/websocket"

type Client struct {
	Id               int
	WebsocketConnect *websocket.Conn
	Hub              *Hub
	Outbound         chan []byte
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		Hub:              hub,
		WebsocketConnect: socket,
		Outbound:         make(chan []byte),
	}
}

func (c *Client) Write() {
	for message := range c.Outbound {
		c.WebsocketConnect.WriteMessage(websocket.TextMessage, message)
	}
}

func (c *Client) Close() {
	c.WebsocketConnect.Close()
	close(c.Outbound)
}
