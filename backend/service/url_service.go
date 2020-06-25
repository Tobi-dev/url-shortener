package service

import (
	"../dto"
	"../model"
	"../repository"
	"github.com/google/uuid"
	"log"
)

func CreateUrl(url *model.Url, repository repository.UrlRepository) dto.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatal(err)
	}

	url.ID = uuidResult.String()

	operationResult := repository.Save(url)

	if operationResult.Error != nil {
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Url)

	return dto.Response{Success: true, Data : data}
}

func FindOneUrlBySlug(slug string, repository repository.UrlRepository) dto.Response{
	operationResult := repository.FindOneById(slug)

	if operationResult.Error != nil{
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Url)

	return dto.Response{Success: true, Data: data}
}

