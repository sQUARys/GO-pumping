package main

import (
	//services "github.com/sQUARys/GO-pumping/blob/dev/Microservices/app/services"
	"microservice/app/providers"
	"microservice/app/repositories"
	"microservice/app/services"
)

func main() {
	provider := providers.New()
	repository := repositories.New()
	service := services.New(provider, repository)

	service.Start()
}
