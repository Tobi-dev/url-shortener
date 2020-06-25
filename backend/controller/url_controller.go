package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../model"
	"../helpers"
	"../service"
	"../repository"
)

func InitUrlRoutes(route *gin.Engine, urlRepository *repository.UrlRepository){
	route.POST("/create", func(context *gin.Context){
		var url model.Url

		err := context.ShouldBindJSON(&url)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := service.CreateUrl(&url, *urlRepository)

		service.CreateUrl(&url, *urlRepository)

		if !response.Success{
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/show/:slug", func(context *gin.Context){
		slug := context.Param("slug")

		code := http.StatusOK

		response := service.FindOneUrlBySlug(slug, *urlRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
