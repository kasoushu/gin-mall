package model

//create commodity

type Address struct {
	Id            uint   `json:"id"`
	UserId        string `json:"user_id"`
	Name          string `json:"name"`
	PostalCode    string `json:"postal_code"`
	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	DetailAddress string `json:"detail_address"`
	IsDefault     bool   `json:"is_default"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
}

type AddressTransfer struct {
	Id            uint     `json:"id"`
	User          UserInfo `json:"user"`
	Name          string   `json:"name"`
	PostalCode    string   `json:"postal_code"`
	Province      string   `json:"province"`
	City          string   `json:"city"`
	District      string   `json:"district"`
	DetailAddress string   `json:"detail_address"`
	IsDefault     bool     `json:"is_default"`
}
