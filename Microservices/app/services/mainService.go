package services

import (
	"microservice/app/providers"
	"microservice/app/repositories"
)

type Service struct {
	Prov providers.Provider
	Repo repositories.Repository
}

func New(provider providers.Provider, repository repositories.Repository) *Service {
	serv := Service{
		Prov: provider,
		Repo: repository,
	}
	return &serv
}

func (serv *Service) GetBodyFromServer() {
	bodyJSON := serv.Prov.GetBodyRequest()
	serv.Prov.UnMarshal(bodyJSON)
}

func (serv *Service) AddToDB(order repositories.Order) {
	serv.Repo.Add(order)
}
