package models

import "strconv"

type Item struct {
	ShortDescription string `json:"shortDescription"`
	PriceStr         string `json:"price"`
}

func (i *Item) Price() (float64, error) {
	return strconv.ParseFloat(i.PriceStr, 64)
}
