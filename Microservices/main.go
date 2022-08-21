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

	controller := controller.New(service)

	r := mux.NewRouter()
	r.HandleFunc("/order/{id}", controller.ReadOrdersId).Methods("POST")
	http.ListenAndServe(":8080", r)
}
