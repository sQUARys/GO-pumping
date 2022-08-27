package providers

import (
	"encoding/json"
	"github.com/sQUARys/GO-pumping/app/order"
	"io"
	"net/http"
)

type Provider struct {
	Url string
}

type orderDTO struct {
	Orders []order.Order `json:"content"`
}

func New() *Provider {
	prov := Provider{
		Url: "http://localhost:8081",
	}
	return &prov
}

func (prov *Provider) GetOrders() ([]order.Order, error) {
	resp, err := http.Get(prov.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var content orderDTO
	err = json.Unmarshal(body, &content)
	if err != nil {
		return nil, err
	}

	return content.Orders, nil
}
