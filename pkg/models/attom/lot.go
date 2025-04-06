package models

type Lot struct {
	LotNum   string  `json:"lotnum"`
	LotSize1 float64 `json:"lotSize1,lotsize1"`
	LotSize2 float64 `json:"lotSize2"`
	PoolType string  `json:"pooltype"`
}
