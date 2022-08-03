package models

type Content struct {
	Content []Data `json:"content"`
}

type Data struct {
	OrderId     int    `json:"order_id"`
	Status      string `json:"status"`
	StoreId     int    `json:"store_id"`
	DateCreated string `json:"date_created"`
}
