package web_socket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/salesforceanton/meower/internal/logger"
	"github.com/salesforceanton/meower/internal/utils"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	Clients    map[int]*Client
	NextId     int
	Register   chan *Client
	Unregister chan *Client
	Mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int]*Client, 0),
		NextId:     0,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Mutex:      new(sync.Mutex),
	}
}

func (h *Hub) Broadcast(message interface{}) {
	messageData, _ := json.Marshal(message)
	for _, client := range h.Clients {
		client.Outbound <- messageData
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.OnConnect(client)

		case client := <-h.Unregister:
			h.OnDisconnect(client)
		}
	}
}

func (h *Hub) OnConnect(client *Client) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	client.Id = h.NextId
	h.NextId++

	h.Clients[client.Id] = client
	log.Println("client connected: ", client.WebsocketConnect.RemoteAddr())
}

func (h *Hub) OnDisconnect(client *Client) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	client.Close()
	delete(h.Clients, client.Id)
	log.Println("client disconnected: ", client.WebsocketConnect.RemoteAddr())
}

func (h *Hub) HandleWebSocket(ctx *gin.Context) {
	socket, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogError(err.Error(), "[Pusher handler]: WebSocket Error")
		return
	}
	client := NewClient(h, socket)
	h.Register <- client

	go client.Write()
}
