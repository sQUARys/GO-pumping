package services

import (
	"github.com/sQUARys/GO-pumping/app/model"
	"time"
)

type Service struct {
	Prov provider
	Repo repositoryOfOrders
}

type provider interface {
	GetOrders() []model.Order
}

type repositoryOfOrders interface {
	Add([]model.Order)
}

func New(provider provider, repository repositoryOfOrders) *Service {
	serv := Service{
		Prov: provider,
		Repo: repository,
	}
	return &serv
}

func (serv *Service) Start() {
	ticker := time.NewTicker(2 * time.Second)

	var orders []model.Order

	for range ticker.C {
		for _, val := range serv.Prov.GetOrders() {
			orders = append(orders, val)
		}
		serv.Repo.Add(orders)
	}

}
