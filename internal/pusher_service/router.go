package pusher_service

import (
	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/eventbus"
	"github.com/salesforceanton/meower/internal/web_socket"
)

type Handler struct {
	eventbus eventbus.EventBus
	hub      *web_socket.Hub
}

func NewHandler(eventbus eventbus.EventBus, hub *web_socket.Hub) *Handler {
	return &Handler{
		eventbus: eventbus,
		hub:      hub,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/pusher", h.hub.HandleWebSocket)

	return router
}
