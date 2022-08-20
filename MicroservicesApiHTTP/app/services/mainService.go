package services

import "github.com/sQUARys/GO-pumping/MicroservicesApiHTTP/app/model"

type repositoryOfOrders interface {
	GetOrdersById(id int) (model.Order, error)
}

type Service struct {
	Repo repositoryOfOrders
}

func New(repo repositoryOfOrders) *Service {
	return &Service{
		Repo: repo,
	}
}

func (serv *Service) StartService(id int) (model.Order, error) {
	order, err := serv.Repo.GetOrdersById(id)
	return order, err
}
