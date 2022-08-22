package services

import (
	"github.com/sQUARys/GO-pumping/app/model"
	"log"
	"time"
)

type Service struct {
	Prov provider
	Repo ordersRepository
}

type provider interface {
	GetOrders() ([]model.Order, error)
}

type ordersRepository interface {
	AddOrders([]model.Order) error
	GetOrderById(id int) (model.Order, error)
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
		orders := serv.GetOrders()
		serv.AddOrders(orders)
	}
}

func (serv *Service) GetOrders() []model.Order {
	orders, err := serv.Prov.GetOrders()
	if err != nil {
		log.Println("Error in service level: ", err)
		return nil
	}
	return orders
}

func (serv *Service) GetOrderById(id int) (model.Order, error) {
	order, err := serv.Repo.GetOrderById(id)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (serv *Service) AddOrders(orders []model.Order) {
	err := serv.Repo.AddOrders(orders)
	if err != nil {
		log.Println("Error in service level: ", err)
		return
	}
}
