package models

type Content struct {
	Content []Order `json:"content"`
}

type Order struct {
	OrderId     int    `json:"order_id"`
	Status      string `json:"status"`
	StoreId     int    `json:"store_id"`
	DateCreated string `json:"date_created"`
}
