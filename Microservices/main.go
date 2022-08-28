package main

import (
	"log"
	"net/http"
	"time"

	controller "github.com/sQUARys/GO-pumping/app/controllers"
	"github.com/sQUARys/GO-pumping/app/providers"
	"github.com/sQUARys/GO-pumping/app/repositories"
	"github.com/sQUARys/GO-pumping/app/routers"
	"github.com/sQUARys/GO-pumping/app/services"
)

func main() {
	provider := providers.New()
	repository := repositories.New()
	service := services.New(provider, repository)

	go service.Start()

	ctr := controller.New(service)
	router := routers.New(ctr)

	router.SetRoutes()

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
		Handler:      router.Router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Error in main : ", err)
	}
}
