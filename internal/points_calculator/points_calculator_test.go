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
