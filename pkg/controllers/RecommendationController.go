package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-search-recommendation-service/internal/clients"
	"net/http"
)

type RecommendationController struct {
	openSearch clients.OpenSearchClient
}

func NewRecommendationController(openSearch clients.OpenSearchClient) *SearchController {
	controller := &SearchController{openSearch}
	return controller
}

func (controller *SearchController) GetRecommendations(context *gin.Context) {
	result, openSearchError := controller.openSearch.QueryRandom()
	if openSearchError != nil {
		context.JSON(http.StatusInternalServerError, openSearchError.Error())
	}

	context.JSON(http.StatusOK, result)
}
