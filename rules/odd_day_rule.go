package rules

// 6 points if the day in the purchase date is odd
type OddDayRule struct{}

func (r OddDayRule) CalculatePoints(receipt Receipt) int64 {
	if receipt.PurchaseDateTime.Day()%2 == 1 {
		return 6
	}
	return 0
}
