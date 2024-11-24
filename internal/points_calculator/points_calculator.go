package points_calculator

import (
	"math"
	"strings"
	"time"

	"github.com/kaiiorg/receipt-processor/internal/models"
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
	// Note: carefully ignoring floating point rounding issues here. If this were more than just a little demo application,
	// we'd want to be way, way more careful and use a proper type designed to handle money.

	quarters := receipt.Total / 0.25
	if quarters == math.Floor(quarters) {
		return 25
	}
	return 0
}

// rulePointPerTwoItems returns 1 point for every two items
// Example: 2 items = 1 point, 4 items = 2 points, 3 items = 1 point, 1 item = 0 points
func (c *Calculator) rulePointPerTwoItems(receipt models.Receipt) uint64 {
	return uint64(len(receipt.Items) / 2)
}

// ruleItemDescriptionMultipleOf3 returns a variable number points if the trimmed length of the string is divisible by 3.
// If it is, the points returned is the price of the item is multiplied by 0.2 and rounded up to the nearest integer
// Example: "Diet Dr. Pepper 2 Liters" @ $2.35 = 1 point,
func (c *Calculator) ruleItemDescriptionMultipleOf3(receipt models.Receipt) uint64 {
	total := uint64(0)

	for _, item := range receipt.Items {
		// Skip if the trimmed length of the description is not divisible by 3
		if len(strings.TrimSpace(item.ShortDescription))%3 != 0 {
			continue
		}

		initialPoints := item.Price * 0.2

		// math.Trunc returns the integer portion of the float: 2.4 -> 2; 2 -> 2
		intPoints := math.Trunc(initialPoints)
		// If the intPoints and initialPoints values are not the same, that means we have a decimal value and we
		// need to add one to round up to the next integer
		if intPoints != initialPoints {
			intPoints++
		}

		total += uint64(intPoints)
	}

	return total
}

// ruleOddDay returns 6 points if the purchase day in the month is an odd number
func (c *Calculator) ruleOddDay(receipt models.Receipt) uint64 {
	date, err := receipt.PurchaseDate()
	// We're going to assume that this receipt has been validated already, but we'll sanity check the error anyway.
	// Invalid receipts don't earn you points.
	if err != nil {
		return 0
	}

	// Another sanity check that the day is within range
	if date.Day() <= 0 || date.Day() > 31 {
		return 0
	}

	if date.Day()%2 != 0 {
		return 6
	}

	return 0
}

// ruleBetweenTimes returns 10 points if the time of purchase is between 2:00PM and 4:00PM noninclusive
func (c *Calculator) ruleBetweenTimes(receipt models.Receipt) uint64 {
	purchaseTime, err := receipt.PurchaseTime()
	// We're going to assume that this receipt has been validated already, but we'll sanity check the error anyway.
	// Invalid receipts don't earn you points.
	if err != nil {
		return 0
	}

	min := time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)
	max := time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)

	if purchaseTime.After(min) && purchaseTime.Before(max) {
		return 10
	}

	return 0
}
