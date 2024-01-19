package query_service

import (
	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/eventbus"
	"github.com/salesforceanton/meower/internal/repository"
	search_repo "github.com/salesforceanton/meower/internal/search"
)

type Handler struct {
	repo       repository.Repository
	eventbus   eventbus.EventBus
	searchRepo search_repo.Repository
}

func NewHandler(
	repo repository.Repository,
	eventbus eventbus.EventBus,
	searchRepo search_repo.Repository) *Handler {
	return &Handler{
		repo:       repo,
		eventbus:   eventbus,
		searchRepo: searchRepo,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/meows", h.corsMiddleware, h.getMeowsListHandler)
	router.GET("/search", h.corsMiddleware, h.searchMeowsHandler)

	return router
}
