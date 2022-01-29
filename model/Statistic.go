package model

type ProductStatisticByStatus struct {
	Total int               `json:"total"`
	List  []ProductByStatus `json:"list"`
}

type ProductByStatus struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type OrderStatistic struct {
	OrderTotal   int `json:"order_total"`
	DayTotal     int `json:"day_total"`
	MonthTotal   int `json:"month_total"`
	DayDone      int `json:"day_done"`
	DayNotDone   int `json:"day_not_done"`
	MonthDone    int `json:"month_done"`
	MonthNotDone int `json:"month_not_done"`
}

type OrderCount struct {
	Day string `json:"day"`
	Cnt int    `json:"cnt"`
}
