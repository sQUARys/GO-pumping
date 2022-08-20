package services

import (
	"github.com/sQUARys/GO-pumping/app/model"
	"log"
	"time"
)

type Service struct {
	Prov provider
	Repo repositoryOfOrders
}

type provider interface {
	GetOrders() ([]model.Order, error)
}

type repositoryOfOrders interface {
	Add([]model.Order) error
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

	for range ticker.C {
		orders, err := serv.Prov.GetOrders()
		if err != nil {
			log.Println("Error : ", err)
			return
		}

		err = serv.Repo.Add(orders)
		if err != nil {
			log.Println("Error : ", err)
			return
		}
	}

}
