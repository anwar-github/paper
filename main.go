package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"papper/controller"
	"papper/database/mysql"
	"papper/repository"
	"papper/service"
	"papper/transformer"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// TODO: move dependency injection to google wire
	newMysql := mysql.NewMysql()
	newRepository := repository.NewRepository(newMysql)
	newService := service.NewService(newRepository)
	newTransformer := transformer.NewTransformer()
	newController := controller.NewController(newService, newTransformer)

	// TODO: move to route
	http.HandleFunc("/", newController.Disbursement)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
