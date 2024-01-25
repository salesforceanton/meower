package query_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/logger"
	"github.com/salesforceanton/meower/internal/utils"
)

type GetMeowsListRequest struct {
	Skip int64 `json:"skip" form:"skip"`
	Take int64 `json:"take" form:"take"`
}

type SearchMeowsRequest struct {
	Skip  int64  `json:"skip" form:"skip"`
	Take  int64  `json:"take" form:"take"`
	Query string `json:"query" form:"query"`
}

func (h *Handler) getMeowsListHandler(ctx *gin.Context) {
	var request GetMeowsListRequest

	if err := ctx.Bind(&request); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		logger.LogError(err.Error(), "[Get Meows List Handler]: Parse request")
		return
	}

	result, err := h.repo.GetMeowsList(ctx.Request.Context(), request.Skip, request.Take)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogError(err.Error(), "[Get Meows List Handler]: Get list from Database")
		return
	}

	utils.NewSuccessResponce(ctx, result)
}

func (h *Handler) searchMeowsHandler(ctx *gin.Context) {
	var request SearchMeowsRequest

	if err := ctx.Bind(&request); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		logger.LogError(err.Error(), "[Search Meows Handler]: Parse request")
		return
	}

	result, err := h.searchRepo.SearchMeows(ctx.Request.Context(), request.Query, request.Skip, request.Take)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogError(err.Error(), "[Search Meows List Handler]: Search from Elastic")
		return
	}

	utils.NewSuccessResponce(ctx, result)
}

func (h *Handler) corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
}
