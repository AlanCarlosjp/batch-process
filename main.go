package main

import (
	"invoice/process/src/application/controller"
	"invoice/process/src/application/service"
	"log"
)

func main() {
	log.Default().Print("Start job invoicing")
	serviceCtx := service.NewBatchService()
	controller.NewBatchController(serviceCtx).InitRoutes()
	log.Default().Print("End job invoicing")
}
