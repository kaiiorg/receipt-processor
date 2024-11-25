package points_calculator

import (
	"testing"

	"github.com/kaiiorg/receipt-processor/internal/models"

	"github.com/stretchr/testify/require"
)

func TestCalculator_Calculate(t *testing.T) {
	// Arrange
	cases := []testCase{
		{
			Receipt: models.Receipt{
				Retailer:        "Target",
				PurchaseDateStr: "2022-01-01",
				PurchaseTimeStr: "13:01",
				Items: []models.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						PriceStr:         "6.49",
					},
					{
						ShortDescription: "Emils Cheese Pizza",
						PriceStr:         "12.25",
					},
					{
						ShortDescription: "Knorr Creamy Chicken",
						PriceStr:         "1.26",
					},
					{
						ShortDescription: "Doritos Nacho Cheese",
						PriceStr:         "3.35",
					},
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						PriceStr:         "12.00",
					},
				},
				TotalStr: "35.35",
			},
			Expected: 28,
		},
		{
			Receipt: models.Receipt{
				Retailer:        "M&M Corner Market",
				PurchaseDateStr: "2022-03-20",
				PurchaseTimeStr: "14:33",
				Items: []models.Item{
					{
						ShortDescription: "Gatorade",
						PriceStr:         "2.25",
					},
					{
						ShortDescription: "Gatorade",
						PriceStr:         "2.25",
					},
					{
						ShortDescription: "Gatorade",
						PriceStr:         "2.25",
					},
					{
						ShortDescription: "Gatorade",
						PriceStr:         "2.25",
					},
				},
				TotalStr: "9.00",
			},
			Expected: 109,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.Calculate(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "receipt from retailer \"%s\"", c.Receipt.Retailer)
	}
}
