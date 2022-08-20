package main

import (
	"github.com/sQUARys/GO-pumping/MicroservicesApiHTTP/app/controllers"
	"github.com/sQUARys/GO-pumping/MicroservicesApiHTTP/app/repositories"
	"github.com/sQUARys/GO-pumping/MicroservicesApiHTTP/app/services"
)

func main() {
	repository := repositories.New()
	service := services.New(repository)
	controller := controller.New(service)
	controller.ReadOrders()
}
