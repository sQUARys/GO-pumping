package main

import (
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
