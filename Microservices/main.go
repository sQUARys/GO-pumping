package main

import (
	"microservice/app/providers"
	"microservice/app/repositories"
	"microservice/app/services"
	"time"
)

func main() {
	provider := providers.New()
	repository := repositories.New()
	service := services.New(*provider, *repository)

	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		service.GetBodyFromServer()
		for i := 0; i < len(service.Prov.Content); i++ {
			service.AddToDB(service.Prov.Content[i])
		}
	}
}
