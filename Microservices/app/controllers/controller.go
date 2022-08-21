package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sQUARys/GO-pumping/app/model"
	"log"
	"net/http"
	"strconv"
)

type Service interface {
	Start()
	GetOrder(int) (model.Order, error)
}

type Controller struct {
	Service Service
}

func New(service Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (ctr *Controller) ReadOrdersById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	idString := vars["id"]

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error strconv in controller level : ", err)
	}

	order, err := ctr.Service.GetOrder(idInt)
	if err != nil {
		log.Println("Error GetOrder in controller level : ", err)
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println("Error json in controller level : ", err)
	}

	w.Write(orderJSON)

}
