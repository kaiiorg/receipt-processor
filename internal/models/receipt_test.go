package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReceipt_PurchaseDate_Valid(t *testing.T) {
	// Arrange
	r := Receipt{
		PurchaseDateStr: "2022-01-01",
	}
	expected, err := time.Parse(time.DateOnly, r.PurchaseDateStr)
	require.NoError(t, err)

	// Act
	result, err := r.PurchaseDate()

	// Assert
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestReceipt_PurchaseDate_Invalid(t *testing.T) {
	// Arrange
	r := Receipt{
		PurchaseDateStr: "this is not a valid date",
	}

	// Act
	result, err := r.PurchaseDate()

	// Assert
	require.ErrorContains(t, err, "cannot parse")
	require.True(t, result.IsZero())
}

func TestReceipt_PurchaseTimeFormat_Valid(t *testing.T) {
	// Arrange
	r := Receipt{
		PurchaseDateTimeStr: "13:01",
	}
	expected, err := time.Parse("15:04", r.PurchaseDateTimeStr)
	require.NoError(t, err)

	// Act
	result, err := r.PurchaseTime()

	// Assert
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestReceipt_PurchaseTimeFormat_Invalid(t *testing.T) {
	// Arrange
	r := Receipt{
		PurchaseDateTimeStr: "this is not a valid date",
	}

	// Act
	result, err := r.PurchaseTime()

	// Assert
	require.ErrorContains(t, err, "cannot parse")
	require.True(t, result.IsZero())
}
