package models

import "time"

const (
	PurchaseDateFormat = time.DateOnly
	PurchaseTimeFormat = "15:04"
)

type Receipt struct {
	Retailer        string  `json:"retailer"`
	PurchaseDateStr string  `json:"purchaseDate"`
	PurchaseTimeStr string  `json:"purchaseDateTime"`
	Total           float64 `json:"total"`
	Items           []Item  `json:"items"`
}

func (r *Receipt) PurchaseDate() (time.Time, error) {
	return time.Parse(PurchaseDateFormat, r.PurchaseDateStr)
}

func (r *Receipt) PurchaseTime() (time.Time, error) {
	return time.Parse(PurchaseTimeFormat, r.PurchaseTimeStr)
}
