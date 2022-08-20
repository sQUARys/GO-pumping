package controller

import (
	"github.com/gorilla/mux"
	"github.com/sQUARys/GO-pumping/MicroservicesApiHTTP/app/model"
	"net/http"
	"strconv"
)

type Service interface {
	StartService(int) (model.Order, error)
}

type Controller struct {
	Serv Service
}

func New(service Service) *Controller {
	return &Controller{
		Serv: service,
	}
}

func (ctr *Controller) ReadOrders() {
	r := mux.NewRouter()

	r.HandleFunc("/post_orders_id/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		idNumber := strconv.Atoi(id)

		ctr.Serv.StartService(idNumber)
	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
