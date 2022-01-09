package model

type Product struct {
	ProductId        uint64  `json:"product_id"`
	CategoryId       uint64  `json:"category_id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Price            uint    `json:"price"`
	Amount           uint    `json:"amount"`
	MainImage        string  `json:"mainImage"`
	Delivery         string  `json:"delivery"`
	Assurance        string  `json:"assurance"`
	Name             string  `json:"name"`
	Weight           float32 `json:"weight"`
	Brand            string  `json:"brand"`
	Origin           string  `json:"origin"`
	ShelfLIfe        uint    `json:"shelf_life"`
	NetWeight        uint    `json:"net_weight"`
	UseWay           string  `json:"use_way"`
	PackingWay       string  `json:"packing_way"`
	StorageCondition string  `json:"storage_condition"`
	DetailImage      string  `json:"detail_image"`
	Status           uint    `json:"status"`
	Created          string  `json:"created"`
	Updated          string  `json:"updated"`
}
