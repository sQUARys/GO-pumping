package services

import "Microservices/providers"

type service struct {
	prov providers.Provider
}

func New(prov provider) *service {
	serv := service{}
}
