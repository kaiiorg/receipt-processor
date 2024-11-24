package points_calculator

import (
	"testing"

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
				Total: 123.0,
			},
			Expected: 50,
		},
		{
			Receipt: models.Receipt{
				Total: 123.32,
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Total: 5.99,
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Total: 5.01,
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Total: 5.0,
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
				Total: 123.0,
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				Total: 123.32,
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Total: 5.99,
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Total: 5.01,
			},
			Expected: 0,
		},
		{
			Receipt: models.Receipt{
				Total: 5.0,
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				Total: 5.25,
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				Total: 5.50,
			},
			Expected: 25,
		},
		{
			Receipt: models.Receipt{
				Total: 5.75,
			},
			Expected: 25,
		},
	}

	for _, c := range cases {
		calc := Calculator{}

		// Act
		result := calc.ruleTotalMultipleOfQuarter(c.Receipt)

		// Assert
		require.Equalf(t, c.Expected, result, "Total: %f", c.Receipt.Total)
	}
}
