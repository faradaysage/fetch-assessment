package rules

import (
	"strings"
)

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up
type ItemDescriptionRule struct{}

func (r ItemDescriptionRule) CalculatePoints(receipt Receipt) int64 {
	var points int64 = 0
	for _, item := range receipt.Items {
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLen%3 == 0 {
			// multiply the price (dollars) by 0.2 (mulitpler) and round up
			// dollars = price / 100
			// multiplier = 20 / 100
			// therefore, factor = (price * 20) / (100 * 100) = (price * 20) / 10000
			// then we "round up" which i interpret as a ceiling function, so we perform a ceiling on our integer
			// from: https://stackoverflow.com/a/2745086 we have q = (x + y - 1) / y, since overflow is not a concern
			// rounded up dollar = ((price * 20) + 10000 - 1) / 10000
			points += int64((item.Price*20)+10000-1) / 10000
		}
	}
	return points
}
