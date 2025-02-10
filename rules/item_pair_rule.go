package rules

// 5 points for every two items on the receipt.
type ItemPairRule struct{}

func (r ItemPairRule) CalculatePoints(receipt Receipt) int64 {
	return int64((len(receipt.Items) / 2) * 5)
}
