package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-search-recommendation-service/internal/clients"
	"net/http"
)

type SearchController struct {
	openSearch clients.OpenSearchClient
}

func NewSearchController(openSearch clients.OpenSearchClient) *SearchController {
	controller := &SearchController{openSearch}
	return controller
}

func (controller *SearchController) SearchVideos(context *gin.Context) {
	query := context.Query("query")
	result, openSearchError := controller.openSearch.SearchVideoDataContaining(query)
	if openSearchError != nil {
		context.JSON(http.StatusInternalServerError, openSearchError.Error())
	}

	context.JSON(http.StatusOK, result)
}
