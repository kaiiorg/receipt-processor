package points_calculator

import (
	"github.com/kaiiorg/receipt-processor/internal/models"
	"math"
)

type Calculator struct{}

func (c *Calculator) Calculate(receipt models.Receipt) uint64 {
	total := c.ruleRetailerName(receipt)
	total += c.ruleRoundTotal(receipt)
	total += c.ruleTotalMultipleOfQuarter(receipt)
	total += c.rulePointPerTwoItems(receipt)
	total += c.ruleItemDescriptionMultipleOf3(receipt)
	total += c.ruleOddDay(receipt)
	total += c.ruleBetweenTimes(receipt)

	return total
}

// ruleRetailerName returns one point for every alphanumeric character in the retailer name
// Example: Target = 6 points, "M&M Corner Market" = 14 points (& and spaces are not alphanumeric!)
func (c *Calculator) ruleRetailerName(receipt models.Receipt) uint64 {
	count := uint64(0)

	for _, character := range receipt.Retailer {
		switch {
		case
			// a-z
			character >= 'a' && character <= 'z',
			// A-Z
			character >= 'A' && character <= 'Z',
			// 0-9
			character >= '0' && character <= '9':
			count++
		}
	}

	return count
}

// ruleRoundTotal returns 50 points if the total is a round dollar amount with no cents
// Example: $7.00 = 50 points, $37.50 = 0 points, $1.23 = 0 points
func (c *Calculator) ruleRoundTotal(receipt models.Receipt) uint64 {
	if receipt.Total == math.Floor(receipt.Total) {
		return 50
	}
	return 0
}

// ruleTotalMultipleOfQuarter returns 25 points if the total is a multiple of $0.25 (1 quarter)
// Example: $7.00 = 25 points, $37.50 = 25 points, $1.23 = 0 points
func (c *Calculator) ruleTotalMultipleOfQuarter(receipt models.Receipt) uint64 {
	return 0
}

// rulePointPerTwoItems returns 1 point for every two items
// Example: 2 items = 1 point, 4 items = 2 points, 3 items = 1 point, 1 item = 0 points
func (c *Calculator) rulePointPerTwoItems(receipt models.Receipt) uint64 {
	return 0
}

// ruleItemDescriptionMultipleOf3 returns a variable number points if the trimmed length of the string is divisible by 3.
// If it is, the points returned is the price of the item is multiplied by 0.2 and rounded up to the nearest integer
// Example: "Diet Dr. Pepper 2 Liters" @ $2.35 = 1 point,
func (c *Calculator) ruleItemDescriptionMultipleOf3(receipt models.Receipt) uint64 {
	return 0
}

// ruleOddDay returns 6 points if the purchase day in the month is an odd number
func (c *Calculator) ruleOddDay(receipt models.Receipt) uint64 {
	return 0
}

// ruleBetweenTimes returns 10 points if the time of purchase is between 2:00PM and 4:00PM noninclusive
func (c *Calculator) ruleBetweenTimes(receipt models.Receipt) uint64 {
	return 0
}
