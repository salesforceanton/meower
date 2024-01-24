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
	router := gin.Default()
	router.GET("/pusher", h.corsMiddleware, h.hub.HandleWebSocket)

	return router
}

func (h *Handler) corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
}
