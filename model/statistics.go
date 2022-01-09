package model

type DataOverviewInfo struct {
	GoodsCnt int     `json:"goods_cnt"`
	OrderCnt int     `json:"order_cnt"`
	Amount   float32 `json:"amount"`
}
