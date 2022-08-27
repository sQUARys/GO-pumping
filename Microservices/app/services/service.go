package services

import (
	"github.com/sQUARys/GO-pumping/app/order"
	"log"
	"time"
)

type Service struct {
	Prov provider
	Repo ordersRepository
}

type provider interface {
	GetOrders() ([]order.Order, error)
}

type ordersRepository interface {
	AddOrders([]order.Order) error
	GetOrderById(id int) (order.Order, error)
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
		serv.Execute()
	}
}

func (serv *Service) Execute() {
	orders , err := serv.GetOrders()
	if err != nil{
		log.Println("Error : " , err)
		return
	}

	err = serv.AddOrders(orders)
	if err != nil{
		log.Println("Error : " , err)
		return
	}
}

func (serv *Service) GetOrders() ([]order.Order , error) {
	orders, err := serv.Prov.GetOrders()
	if err != nil {
		return nil , err
	}
	return orders , nil
}

func (serv *Service) GetOrderById(id int) (order.Order, error) {
	o, err := serv.Repo.GetOrderById(id)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}

func (serv *Service) AddOrders(orders []order.Order ) error {
	err := serv.Repo.AddOrders(orders)
	if err != nil {
		return err
	}
	return nil
}
