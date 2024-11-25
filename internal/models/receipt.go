package models

import (
	"strconv"
	"time"
)

const (
	PurchaseDateFormat = time.DateOnly
	PurchaseTimeFormat = "15:04"
)

type Receipt struct {
	Retailer        string `json:"retailer"`
	PurchaseDateStr string `json:"purchaseDate"`
	PurchaseTimeStr string `json:"purchaseDateTime"`
	TotalStr        string `json:"total"`
	Items           []Item `json:"items"`
}

func (r *Receipt) PurchaseDate() (time.Time, error) {
	return time.Parse(PurchaseDateFormat, r.PurchaseDateStr)
}

func (r *Receipt) PurchaseTime() (time.Time, error) {
	return time.Parse(PurchaseTimeFormat, r.PurchaseTimeStr)
}

func (r *Receipt) Total() (float64, error) {
	return strconv.ParseFloat(r.TotalStr, 64)
}
