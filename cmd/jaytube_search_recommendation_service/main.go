package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-search-recommendation-service/internal/clients"
	"github.com/the-jay-team/jaytube-search-recommendation-service/pkg/configs"
	"github.com/the-jay-team/jaytube-search-recommendation-service/pkg/controllers"
)

func main() {
	server := gin.Default()
	openSearchConfig := configs.GetEnvironmentConfig().OpenSearch
	openSearchClient := *clients.NewOpenSearchClient(
		openSearchConfig.Host,
		openSearchConfig.VideoDataIndex,
		openSearchConfig.Username,
		openSearchConfig.Password)

	searchController := controllers.NewSearchController(openSearchClient)
	recommendationController := controllers.NewRecommendationController(openSearchClient)

	server.GET("/search", searchController.SearchVideos)
	server.GET("/recommendation", recommendationController.GetRecommendations)

	server.Run(":8080")
}
