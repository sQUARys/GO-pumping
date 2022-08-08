package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	dbRepo "microservice/app/repositories"
	"net/http"
)

type Provider struct {
	Content []dbRepo.Order `json:"content"`
}

func New() *Provider {
	prov := Provider{}
	return &prov
}

func (prov *Provider) GetBodyRequest() []byte {
	url := "http://localhost:8081"
	resp, err := http.Get(url)
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

func (prov *Provider) UnMarshal(body []byte) {
	err := json.Unmarshal(body, prov)
	if err != nil {
		log.Println("Error: ", err)
	}
}
