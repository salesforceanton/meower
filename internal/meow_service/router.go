package meow_service

import (
	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/eventbus"
	"github.com/salesforceanton/meower/internal/repository"
)

type Handler struct {
	repo     repository.Repository
	eventBus eventbus.EventBus
}

func NewHandler(repo repository.Repository,
	eventBus eventbus.EventBus) *Handler {
	return &Handler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.POST("/meows", h.createMeowHandler)
	router.OPTIONS("/meows", h.preflightHandler)

	return router
}
