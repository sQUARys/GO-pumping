package providers

import (
	"encoding/json"
	"github.com/sQUARys/GO-pumping/app/model"
	"io/ioutil"
	"log"
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

func (prov *Provider) GetLocalhostBodyRequest() []model.Order {
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

	var content model.Content
	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Println("Error: ", err)
	}

	return content.Orders
}
