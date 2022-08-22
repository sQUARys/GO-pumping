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
	Add([]model.Order) error
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
		serv.AddOrders()
	}
}

func (serv *Service) AddOrders() {
	order, err := serv.Prov.GetOrders()
	if err != nil {
		log.Println("Error in service level: ", err)
		return
	}

	err = serv.Repo.Add(order)
	if err != nil {
		log.Println("Error in service level: ", err)
		return
	}
}

func (serv *Service) GetOrder(id int) (model.Order, error) {
	order, err := serv.Repo.GetOrderById(id)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
