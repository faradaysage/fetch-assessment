package rules

// 25 points if the total is a multiple of 0.25.
type MultipleOfQuarterRule struct{}

func (r MultipleOfQuarterRule) CalculatePoints(receipt Receipt) int64 {
	// we're storing the value in cents, so $0.25 -> 25 cents
	if receipt.Total%25 == 0 {
		return 25
	}
	return 0
}
