package models

import "errors"

var (
	ErrInvalidDate            = errors.New("invalid date")
	ErrInvalidTime            = errors.New("invalid time")
	ErrInvalidTotal           = errors.New("invalid total")
	ErrInvalidPrice           = errors.New("invalid price")
	ErrItemsDoNotAddUpToTotal = errors.New("items do not add up to total")
)
