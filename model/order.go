package model

type Order struct {
	Id         uint64  `json:"id"`
	ProductId  uint64  `json:"product_id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	AddressId  uint64  `json:"address_id"`
	UserId     string  `json:"user_id"`
	Created    string  `json:"created"`
	Updated    string  `json:"updated"`
	AdminId    string  `json:"admin_id"`
	Amount     int     `json:"amount"`
}

// to show on admin
type OrderTransfer struct {
	Id          uint64          `json:"id"`
	ProductName string          `json:"product_name"`
	Address     AddressTransfer `json:"address"`
	TotalPrice  float64         `json:"total_price"`
	Status      string          `json:"status"`
	UserInfo    UserInfo        `json:"user_info"`
	Created     string          `json:"created"`
	Updated     string          `json:"updated"`
	Amount      int             `json:"amount"`
}
