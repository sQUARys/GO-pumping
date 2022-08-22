package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

//И еще. Из контроллера и роутера убери интерфейсы. И передавай пря по типу.

type Ctr interface {
	GetOrderById(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	Router     *mux.Router
	Controller Ctr
}

func New(controller Ctr) *Router {
	r := mux.NewRouter()
	return &Router{
		Controller: controller,
		Router:     r,
	}
}

func (r *Router) SetRoutes() {
	r.Router.HandleFunc("/order/{id}", r.Controller.GetOrderById).Methods("POST")
}
