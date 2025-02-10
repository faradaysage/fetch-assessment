package rules

// 50 points if the total is a round dollar amount with no cents.
type RoundDollarRule struct{}

func (r RoundDollarRule) CalculatePoints(receipt Receipt) int64 {
	if receipt.Total%100 == 0 {
		return 50
	}
	return 0
}
