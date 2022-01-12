package model

type Order struct {
	Id          uint64  `json:"id"`
	ProductItem string  `json:"product_item"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
	AddressId   uint64  `json:"address_id"`
	UserId      string  `json:"user_id"`
	NickName    string  `json:"nick_name"`
	Created     string  `json:"created"`
	Updated     string  `json:"updated"`
}
