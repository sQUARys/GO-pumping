package main

import (
	"github.com/gorilla/mux"
	controller "github.com/sQUARys/GO-pumping/app/controllers"
	"github.com/sQUARys/GO-pumping/app/providers"
	"github.com/sQUARys/GO-pumping/app/repositories"
	"github.com/sQUARys/GO-pumping/app/services"
	"net/http"
)

func main() {
	provider := providers.New()
	repository := repositories.New()
	service := services.New(provider, repository)
	go service.Start()

	r := mux.NewRouter()
	controller := controller.New(service)
	controller.ReadOrdersId(r)
	http.ListenAndServe(":8080", r)
}
