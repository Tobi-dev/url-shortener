package main

import (
	"./configs"
	Connection "./database"
	"./model"
	"./repository"
	"fmt"
	"./helpers"
)

func main(){

	fmt.Println(helpers.GenerateRandomSlug())

	db := Connection.GetDB()

	db.AutoMigrate(&model.Url{})

	defer db.Close()

	urlRepository := repository.NewUrlRepository(db)

	route := configs.SetupRoutes(urlRepository)

	route.Run(":8000")
}