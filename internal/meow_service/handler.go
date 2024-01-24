package meow_service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/logger"
	"github.com/salesforceanton/meower/internal/schema"
	"github.com/salesforceanton/meower/internal/utils"
	"github.com/segmentio/ksuid"
)

type CreateMeowRequest struct {
	Body string `json:"body"`
}

func (h *Handler) createMeowHandler(ctx *gin.Context) {
	var request CreateMeowRequest

	// Parse request
	if err := ctx.BindJSON(&request); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		logger.LogError(err.Error(), "[Create Meow Handler]: Parse request")
		return
	}

	// Generate Meow Struct
	created_at := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(created_at)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogError(err.Error(), "[Create Meow Handler]: Generate Id")
		return
	}

	message := schema.Meow{
		Id:        id.String(),
		Body:      request.Body,
		CreatedAt: created_at,
	}

	// Insert Meow Message record in db
	if err = h.repo.InsertMeow(ctx.Request.Context(), message); err != nil {
		utils.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogError(err.Error(), "[Create Meow Handler]: Insert Record in DB")
		return
	}

	// Publish Create Meow Event into Event Bus
	if err = h.eventBus.PublishMeowCreated(message); err != nil {
		utils.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogError(err.Error(), "[Create Meow Handler]: Publish Event Bus message")
		return
	}

	utils.NewSuccessResponce(ctx, "Meow has been created Successfully!")
}

func (h *Handler) preflightHandler(ctx *gin.Context) {
	return
}

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
