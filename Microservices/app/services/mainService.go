package services

import (
	"microservice/app/models"
)

type Service struct {
	prov providerInterface
	repo repositoryInterface
}

type providerInterface interface {
	GetBodyRequest() []byte
	UnMarshal([]byte)
}

type repositoryInterface interface {
	Add(models.Order)
}

func New(provider providerInterface, repository repositoryInterface) *Service {
	serv := Service{
		prov: provider,
		repo: repository,
	}
	return &serv
}

func (serv *Service) GetBodyFromServer() {
	bodyJSON := serv.prov.GetBodyRequest()
	serv.prov.UnMarshal(bodyJSON)
}

func (serv *Service) AddToDB(order models.Order) {
	serv.repo.Add(order)
}
