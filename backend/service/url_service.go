package service

import (
	"../dto"
	"../model"
	"../repository"
	"github.com/google/uuid"
	"log"
	"../helpers"
)

func CreateUrl(url *model.Url, repository repository.UrlRepository) dto.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatal(err)
	}

	//Lets think about worst case
	existsAlready := true;
	tmpSlug := ""
	if(len(url.Slug) <= 0){
		tmpSlug = helpers.GenerateRandomSlug()
	}else{
		tmpSlug = url.Slug
	}

	for existsAlready {
		operationResult := repository.FindOneBySlug(tmpSlug)
		if operationResult.Error != nil {
			url.Slug = tmpSlug
			existsAlready = false
		}else{
			tmpSlug = helpers.GenerateRandomSlug()
		}
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
	operationResult := repository.FindOneBySlug(slug)

	if operationResult.Error != nil{
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Url)

	return dto.Response{Success: true, Data: data}
}

