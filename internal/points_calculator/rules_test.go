package points_calculator

import (
	"testing"
	"time"

	"github.com/kaiiorg/receipt-processor/internal/models"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Receipt  models.Receipt
	Expected uint64
}

func TestCaculator_RuleRetailerName(t *testing.T) {
	// Arrange
	cases := []testCase{
		{
			Receipt: models.Receipt{
				Retailer: "Target",
			},
			Expected: 6,
		},
		{
			Receipt: models.Receipt{
				Retailer: "M&M Corner Market",
			},
			Expected: 14,
		},
		{
			Receipt: models.Receipt{
				Retailer: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			},
			Expected: 62,
		},
		{
			Receipt: models.Receipt{
				Retailer: "`~!@#$%^&*()_+-=[{]},<.>|\\/?;:'\"☺☻",
			},
			Expected: 0,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.ruleRetailerName(c.Receipt)

		// Assert
		require.Equal(t, c.Expected, result)
	}
}

func TestCalculator_RuleRoundTotal(t *testing.T) {
	// Arrange
	cases := []testCase{
		{
			Receipt: models.Receipt{
				TotalStr: "123.0",
			},
			Expected: 50,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "123.32",
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.99",
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.01",
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.0",
			},
			Expected: 50,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.ruleRoundTotal(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "Total: %f", c.Receipt.Total)
	}
}

func TestCalculator_RuleTotalMultipleOfQuarter(t *testing.T) {
	// Arrange
	cases := []testCase{
		{
			Receipt: models.Receipt{
				TotalStr: "123.0",
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "123.32",
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.99",
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.01",
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.0",
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.25",
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.50",
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				TotalStr: "5.75",
			},
			Expected: 25,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.ruleTotalMultipleOfQuarter(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "Total: %s", c.Receipt.TotalStr)
	}
}

func TestCalculator_RulePointPerTwoItems(t *testing.T) {
	cases := []testCase{
		{
			Receipt: models.Receipt{
				Items: []models.Item{
					// 1
					{},
				},
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Items: []models.Item{
					// 2
					{}, {},
				},
			},
			Expected: 5,
		},
		{
			Receipt: models.Receipt{
				Items: []models.Item{
					// 3
					{}, {}, {},
				},
			},
			Expected: 5,
		},
		{
			Receipt: models.Receipt{
				Items: []models.Item{
					// 4
					{}, {}, {}, {},
				},
			},
			Expected: 10,
		},
		{
			Receipt: models.Receipt{
				Items: []models.Item{
					// 0
				},
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Items: []models.Item{
					// 28
					{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				},
			},
			Expected: 70,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.rulePointPerTwoItems(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "Item count: %d", len(c.Receipt.Items))
	}
}

func TestCalculator_RuleItemDescriptionMultipleOf3(t *testing.T) {
	// Arrange
	cases := []testCase{
		{
			Receipt: models.Receipt{
				Retailer: "1",
				Items: []models.Item{
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						PriceStr:         "12.00",
					},
				},
			},
			Expected: 3,
		},
		{
			Receipt: models.Receipt{
				Retailer: "2",
				Items: []models.Item{
					{
						ShortDescription: "Diet Dr. Pepper 2 Liters",
						PriceStr:         "2.35",
					},
				},
			},
			Expected: 1,
		},
		{
			Receipt: models.Receipt{
				Retailer: "2",
				Items: []models.Item{
					{
						ShortDescription: "Dr. Pepper 2 Liters",
						PriceStr:         "2.35",
					},
				},
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Retailer: "2",
				Items: []models.Item{
					{
						ShortDescription: "Diet Dr. Pepper 2 Liters",
						PriceStr:         "2.35",
					},
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						PriceStr:         "12.00",
					},
					{
						ShortDescription: "Dr. Pepper 2 Liters",
						PriceStr:         "2.35",
					},
				},
			},
			Expected: 4,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.ruleItemDescriptionMultipleOf3(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "Receipt from retailer \"%s\"", c.Receipt.Retailer)
	}
}

func TestCalculator_RuleOddDay(t *testing.T) {
	// Arrange
	cases := []testCase{
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 0, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 1, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 6,
		},
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 2, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 3, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 6,
		},
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 4, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 29, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 6,
		},
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 30, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				PurchaseDateStr: time.Date(2024, 24, 31, 0, 0, 0, 0, time.UTC).Format(time.DateOnly),
			},
			Expected: 6,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.ruleOddDay(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "Date: %s", c.Receipt.PurchaseDateStr)
	}
}

func TestCalculator_RuleBetweenTimes(t *testing.T) {
	// Arrange
	cases := []testCase{
		{
			Receipt: models.Receipt{
				PurchaseTimeStr: time.Date(0, 0, 0, 14, 1, 0, 0, time.UTC).Format("15:04"),
			},
			Expected: 10,
		},
		{
			Receipt: models.Receipt{
				PurchaseTimeStr: time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC).Format("15:04"),
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				PurchaseTimeStr: time.Date(0, 0, 0, 15, 0, 0, 0, time.UTC).Format("15:04"),
			},
			Expected: 10,
		},
		{
			Receipt: models.Receipt{
				PurchaseTimeStr: time.Date(0, 0, 0, 15, 30, 0, 0, time.UTC).Format("15:04"),
			},
			Expected: 10,
		},
		{
			Receipt: models.Receipt{
				PurchaseTimeStr: time.Date(0, 0, 0, 15, 59, 0, 0, time.UTC).Format("15:04"),
			},
			Expected: 10,
		},
		{
			Receipt: models.Receipt{
				PurchaseTimeStr: time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC).Format("15:04"),
			},
			Expected: 0,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.ruleBetweenTimes(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "Time: %s", c.Receipt.PurchaseTimeStr)
	}
}
