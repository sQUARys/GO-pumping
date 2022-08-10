package services

import (
	"microservice/app/model"
	"time"
)

type Service struct {
	Prov provider
	Repo repositoryOfOrders
}

type provider interface {
	GetLocalhostBodyRequest() []model.Order
}

type repositoryOfOrders interface {
	Add(model.Order)
}

func New(provider provider, repository repositoryOfOrders) *Service {
	serv := Service{
		Prov: provider,
		Repo: repository,
	}
	return &serv
}

func (serv *Service) GetOrdersFromServer() []model.Order {
	orders := serv.Prov.GetLocalhostBodyRequest()
	return orders
}

func (serv *Service) Start() {
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		orders := serv.GetOrdersFromServer()
		for i := 0; i < len(orders); i++ {
			serv.Repo.Add(orders[i])
		}
	}
}
