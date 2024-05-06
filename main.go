package main

import (
	"github.com/joho/godotenv"
	"invoice/process/src/application/controller"
	"invoice/process/src/application/service"
	"invoice/process/src/infrastructure/db"
	"log"
)

func main() {
	log.Default().Print("Start job invoicing")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error in loading env: ", err)
		return
	}
	db := db.CreateConnection()
	serviceCtx := service.NewBatchService(db)
	controller.NewBatchController(serviceCtx).InitRoutes()
	log.Default().Print("End job invoicing")
}
