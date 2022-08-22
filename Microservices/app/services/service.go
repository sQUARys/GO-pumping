package services

import (
	"github.com/sQUARys/GO-pumping/app/order"
	"log"
	"time"
)

type Service struct {
	Prov provider
	Repo ordersRepository
}

type provider interface {
	GetOrders() ([]order.Order, error)
}

type ordersRepository interface {
	AddOrders([]order.Order) error
	GetOrderById(id int) (order.Order, error)
}

func New(provider provider, repository ordersRepository) *Service {
	serv := Service{
		Prov: provider,
		Repo: repository,
	}
	return &serv
}

func (serv *Service) Start() {
	ticker := time.NewTicker(2 * time.Second)

	for range ticker.C {
		serv.Execute()
	}
}

func (serv *Service) Execute() {
	orders := serv.GetOrders()
	serv.AddOrders(orders)
}

func (serv *Service) GetOrders() []order.Order {
	orders, err := serv.Prov.GetOrders()
	if err != nil {
		log.Println("Error in service level: ", err)
		return nil
	}
	return orders
}

func (serv *Service) GetOrderById(id int) (order.Order, error) {
	o, err := serv.Repo.GetOrderById(id)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}

func (serv *Service) AddOrders(orders []order.Order) {
	err := serv.Repo.AddOrders(orders)
	if err != nil {
		log.Println("Error in service level: ", err)
		return
	}
}
