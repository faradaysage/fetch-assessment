package rules

import (
	"regexp"
)

// One point for every alphanumeric character in the retailer name
type RetailerNameRule struct{}

func (r RetailerNameRule) CalculatePoints(receipt Receipt) int64 {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return int64(len(re.FindAllString(receipt.Retailer, -1)))
}
