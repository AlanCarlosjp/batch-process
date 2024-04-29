package main

import (
	"invoice/process/src/application/controller"
	"log"
)

func main() {
	log.Default().Print("Start job invoicing")
	controller.NewBatchController().InitRoutes()
	log.Default().Print("End job invoicing")
}
