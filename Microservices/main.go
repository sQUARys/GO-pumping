package main

import (
	"github.com/sQUARys/GO-pumping/app/providers"
	"github.com/sQUARys/GO-pumping/app/repositories"
	"github.com/sQUARys/GO-pumping/app/services"
)

func main() {
	provider := providers.New()
	repository := repositories.New()
	service := services.New(provider, repository)

	service.Start()
}
