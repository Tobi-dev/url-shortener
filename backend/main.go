package main

import (
	Connection "./database"
	"./model"
	"./repository"
	"./configs"
	"./helpers"
)

func main(){

	helpers.GenerateRandomSlug("http://google.de")

	db := Connection.GetDB()

	db.AutoMigrate(&model.Url{})

	defer db.Close()

	urlRepository := repository.NewUrlRepository(db)

	route := configs.SetupRoutes(urlRepository)

	route.Run(":8000")
}