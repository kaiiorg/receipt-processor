package models

import (
	"errors"
	"fmt"
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
	PurchaseTimeStr string `json:"purchaseTime"`
	TotalStr        string `json:"total"`
	Items           []Item `json:"items"`
}

func (r *Receipt) Validate() error {
	_, err := r.PurchaseDate()
	if err != nil {
		return errors.Join(ErrInvalidDate, err)
	}

	_, err = r.PurchaseTime()
	if err != nil {
		return errors.Join(ErrInvalidTime, err)
	}

	expectedTotal, err := r.Total()
	if err != nil {
		return errors.Join(ErrInvalidTotal, err)
	}

	actualTotal := float64(0)
	for i, item := range r.Items {
		price, err := item.Price()
		if err != nil {
			return errors.Join(
				ErrInvalidPrice,
				fmt.Errorf("item %d, \"%s\"", i, item.ShortDescription),
				err,
			)
		}
		actualTotal += price
	}

	if actualTotal != expectedTotal {
		return ErrItemsDoNotAddUpToTotal
	}

	return nil
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
