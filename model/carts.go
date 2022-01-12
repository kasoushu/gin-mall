package model

type Cart struct {
	ProductId    uint   `json:"product_id"`
	ProductCount uint   `json:"product_count"`
	UserId       string `json:"user_id"`
}
