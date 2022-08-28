package routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sQUARys/GO-pumping/app/controllers"
)

type Router struct {
	Router     *mux.Router
	Controller controller.Controller
}

func New(controller *controller.Controller) *Router {
	r := mux.NewRouter()

	return &Router{
		Controller: *controller,
		Router:     r,
	}
}

func (r *Router) SetRoutes() {
	r.Router.HandleFunc("/order/{id}", r.Controller.GetOrderById).Methods("POST")
}
