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
						Price:            6.49,
					},
					{
						ShortDescription: "Emils Cheese Pizza",
						Price:            12.25,
					},
					{
						ShortDescription: "Knorr Creamy Chicken",
						Price:            1.26,
					},
					{
						ShortDescription: "Doritos Nacho Cheese",
						Price:            3.35,
					},
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            12.00,
					},
				},
				Total: 35.35,
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
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
					{
						ShortDescription: "Gatorade",
						Price:            2.25,
					},
				},
				Total: 9.00,
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
