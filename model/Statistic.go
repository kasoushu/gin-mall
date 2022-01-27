package model

type ProductStatisticByStatus struct {
	Total int               `json:"total"`
	List  []ProductByStatus `json:"list"`
}

type ProductByStatus struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}
