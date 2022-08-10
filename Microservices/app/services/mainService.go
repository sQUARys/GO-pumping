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
	GetLocalhostBodyRequest() []byte
	UnMarshalBodyRequest([]byte) []model.Order
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

func (serv *Service) GetBodyFromServer() []model.Order {
	bodyJSON := serv.Prov.GetLocalhostBodyRequest()
	bodyUnMarshalled := serv.Prov.UnMarshalBodyRequest(bodyJSON)
	return bodyUnMarshalled
}

func (serv *Service) SendOrdersToDB() {
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		orders := serv.GetBodyFromServer()
		for i := 0; i < len(orders); i++ {
			serv.Repo.Add(orders[i])
		}
	}
}
