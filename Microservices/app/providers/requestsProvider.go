package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"microservice/app/models"
	"net/http"
)

type Provider struct {
	Url string
}

func New() *Provider {
	prov := Provider{
		Url: "http://localhost:8081",
	}
	return &prov
}

func (prov *Provider) GetLocalhostBodyRequest() []byte {
	resp, err := http.Get(prov.Url)
	if err != nil {
		log.Println("Error: ", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: ", err)
		return nil
	}

	return body
}

func (prov *Provider) UnMarshalBodyRequest(body []byte) []models.Order {
	var content models.Content
	err := json.Unmarshal(body, &content)
	if err != nil {
		log.Println("Error: ", err)
	}
	return content.Orders
}
