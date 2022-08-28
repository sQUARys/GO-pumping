package providers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sQUARys/GO-pumping/app/order"
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

func (prov *Provider) GetOrders() (orders []order.Order, err error) {
	ctx := context.Background()

	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, prov.Url, nil)
	if err != nil {
		return
	}

	defer func() { err = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var content orderDTO

	err = json.Unmarshal(body, &content)
	if err != nil {
		return
	}

	orders = content.Orders

}
