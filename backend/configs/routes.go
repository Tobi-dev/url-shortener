package configs

import (
	"../repository"
	"github.com/gin-gonic/gin"
	urlController "../controller"
)

func SetupRoutes(urlRepository *repository.UrlRepository) *gin.Engine {
	route := gin.Default()
	urlController.InitUrlRoutes(route, urlRepository)

	return route
}