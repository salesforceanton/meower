package query_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/meower/internal/logger"
	"github.com/salesforceanton/meower/internal/utils"
)

type GetMeowsListRequest struct {
	Skip int64 `json:"skip"`
	Take int64 `json:"take"`
}

type SearchMeowsRequest struct {
	Skip  int64  `json:"skip"`
	Take  int64  `json:"take"`
	Query string `json:"query"`
}

func (h *Handler) getMeowsListHandler(ctx *gin.Context) {
	var request GetMeowsListRequest

	if err := ctx.BindJSON(request); err != nil {
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

	if err := ctx.BindJSON(request); err != nil {
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
