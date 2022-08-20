package main

import (
	controller "github.com/sQUARys/GO-pumping/app/controllers"
	"github.com/sQUARys/GO-pumping/app/providers"
	"github.com/sQUARys/GO-pumping/app/repositories"
	"github.com/sQUARys/GO-pumping/app/services"
)

func main() {
	provider := providers.New()
	repository := repositories.New()
	service := services.New(provider, repository)

	go service.Start()

	controller := controller.New(service)
	controller.ReadOrdersId()
}
