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
	Start() error
	GetOrders(int) ([]model.Order, error)
}

type Controller struct {
	Serv Service
}

func New(service Service) *Controller {
	return &Controller{
		Serv: service,
	}
}

func (ctr *Controller) ReadOrdersId() {
	r := mux.NewRouter()

	r.HandleFunc("/order/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]

		idNumber, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error strconv in controller level : ", err)
		}

		orders, err := ctr.Serv.GetOrders(idNumber)
		if err != nil {
			log.Println("Error GetOrders in controller level : ", err)
		}

		ordersJSON, err := json.Marshal(orders)
		if err != nil {
			log.Println("Error json in controller level : ", err)
		}

		w.Write(ordersJSON)

	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
