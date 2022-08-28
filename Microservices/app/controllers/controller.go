package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sQUARys/GO-pumping/app/services"
)

type Controller struct {
	Service services.Service
}

func New(service *services.Service) *Controller {
	return &Controller{
		Service: *service,
	}
}

func (ctr *Controller) GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	idString := vars["id"]

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error strconv in controller level : ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	order, err := ctr.Service.GetOrderById(idInt)
	if err != nil {
		log.Println("Error GetOrder in controller level : ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println("Error json in controller level : ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	status, err := w.Write(orderJSON)
	if err != nil {
		log.Println("Error in writing  level : ", err)
		w.WriteHeader(status)

		return
	}
}
