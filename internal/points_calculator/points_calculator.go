package points_calculator

import (
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
