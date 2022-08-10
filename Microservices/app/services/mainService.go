package services

import (
	"microservice/app/model"
	"time"
)

type Service struct {
	Prov providerInterface
	Repo repositoryInterface
}

type providerInterface interface {
	GetLocalhostBodyRequest() []byte
	UnMarshalBodyRequest([]byte) []model.Order
}

type repositoryInterface interface {
	Add(model.Order)
}

func New(provider providerInterface, repository repositoryInterface) *Service {
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
