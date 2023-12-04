package meow_service

import (
	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/eventbus"
	"github.com/salesforceanton/meower/internal/repository"
)

type Handler struct {
	repo     *repository.PostgresRepo
	eventBus *eventbus.NatsEventbus
}

func NewHandler(repo *repository.PostgresRepo,
	eventBus *eventbus.NatsEventbus) *Handler {
	return &Handler{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/meows", h.createMeowHandler)

	return router
}
